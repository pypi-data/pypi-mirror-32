#ifndef ABSYNTYPES
#define ABSYNTYPES
#ifndef NODESTTYPE
#define NODESTTYPE
typedef long long int key_t;

typedef struct NodeST_s NodeST;

typedef enum {UNIT_N, ID_N, EXP_N, EXPLIST_N, STMT_N, STMTLIST_N, ERROR_N, GRAMMAR_N} NodeType;
#endif

/*
  There is a circular dependency between parser.tab.h (which contains the tokens)
  and this file (since parser.tab.h needs the node structures), so we define the
  struct types before including parser.tab.h, and define the actual structs after.
*/

typedef struct Unit_s Unit;
typedef struct Id_s Id;
typedef struct Expression_s Expression;
typedef struct ExpressionList_s ExpressionList;
typedef struct Statement_s Statement;
typedef struct StatementList_s StatementList;
typedef struct Error_s Error;
typedef struct Grammar_s Grammar;

#ifndef YYTOKENTYPE
# define YYTOKENTYPE
  enum yytokentype
  {
    EPSILON = 258,
    ATOM = 259,
    Z = 260,
    UNION = 261,
    PROD = 262,
    SET = 263,
    POWERSET = 264,
    SEQUENCE = 265,
    CYCLE = 266,
    SUBST = 267,
    CARD = 268,
    LPAR = 269,
    RPAR = 270,
    COMMA = 271,
    LEQ = 272,
    GEQ = 273,
    EQ = 274,
    ID = 275,
    NUMBER = 276
  };
#endif

typedef enum {NONE, LESS, EQUAL, GREATER} Restriction; // restrictions to cardinality

typedef enum {LEXER, PARSER} ErrorType; // origin of error

typedef enum {ISERROR, NOTERROR} GrammarType; // types of grammars resulting from parsing

/***************************** Abstract Syntax Tree Nodes *****************************/

/*
  Node for Atom, Epsilon or Z.
*/
struct Unit_s
{
  enum yytokentype type;
  key_t key;
  char* (*toString)(const struct Unit_s* self);
  char* (*toJson)(const struct Unit_s* self);
};

/*
  Node for Id.
*/
struct Id_s
{
  char* name;
  key_t key;
  char* (*toString)(const struct Id_s* self);
  char* (*toJson)(const struct Id_s* self);
};

/* 
   Node for expression. The component contains the node below the expression in the
   abstract syntax tree, the type indicates what kind of expression this is (Union,
   Prod, Set, Id, Atom, ...), restriction and limit can be used to add restrictions
   to the cardinality.
*/
struct Expression_s
{
  void* component; // could be pointer to Expression, ExpressionList, Id or Unit
  enum yytokentype type;
  Restriction restriction; // restriction type
  long long int limit; // numerical value of restriction in cardinality
  key_t key;
  char* (*toString)(const struct Expression_s* self);
  char* (*toJson)(const struct Expression_s* self);
};

/*
  Node for a list (comma-separated in the input) of expressions.
*/
struct ExpressionList_s
{
  Expression** components;
  int size; // number of expressions in the list of components
  int space; // maximum number of expressions that can be put in the current list
  key_t key;
  char* (*toString)(const struct ExpressionList_s* self);
  char* (*toJson)(const struct ExpressionList_s* self);
};

/*
  Node for a statement. This should be comething of the form a = exp, where
  a is an Id and exp is an expression.
*/
struct Statement_s
{
  Id* variable;
  Expression* expression;
  key_t key;
  char* (*toString)(const struct Statement_s* self);
  char* (*toJson)(const struct Statement_s* self);
};

/*
  Node for a list (comma-separated in the input) of statements.
*/ 
struct StatementList_s
{
  Statement** components;
  int size; // number of statements in the list of components
  int space; // maximum number of statements that can be put in the current list
  key_t key;
  char* (*toString)(const struct StatementList_s* self);
  char* (*toJson)(const struct StatementList_s* self);
};

/*
  Node used to report errors.
*/
struct Error_s
{
  int line;
  char* message;
  ErrorType type;
  key_t key;
  char* (*toString)(const struct Error_s* self);
  char* (*toJson)(const struct Error_s* self);
};

/*
  Root of the abstract syntax tree, can be an error or a list of statements.
*/
struct Grammar_s
{
  GrammarType type;
  void* component; // can be Error or StatementList
  key_t key;
  char* (*toString)(const struct Grammar_s* self);
  char* (*toJson)(const struct Grammar_s* self);
};
#endif
#ifndef C2J_H
#define C2H_H
Grammar* readGrammar(char* filename);
#endif
