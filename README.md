### monkey-interpreter
『Go言語でつくるインタプリタ』の実装コード  

![51sLCPa8DBL _SX388_BO1,204,203,200_](https://user-images.githubusercontent.com/43651940/96365304-9cd7c200-117a-11eb-9a00-701da2df50ae.jpg)


### chapter01 字句解析  
`$ go run main.go`  
Hello kotarokashihara! This is the monkey Programming Language  
Feel free to type commands  
`>> let add = fn(x,y){x+y;};`  
{Type:LET Literal:let}  
{Type:IDENT Literal:add}  
{Type:= Literal:=}  
{Type:FUNCTION Literal:fn}  
{Type:( Literal:(}  
{Type:IDENT Literal:x}  
{Type:, Literal:,}  
{Type:IDENT Literal:y}  
{Type:) Literal:)}  
{Type:{ Literal:{}  
{Type:IDENT Literal:x}  
{Type:+ Literal:+}  
{Type:IDENT Literal:y}  
{Type:; Literal:;}  
{Type:} Literal:}}  
{Type:; Literal:;}  

### chapter02 構文解析  
