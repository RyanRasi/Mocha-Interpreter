package evaluator

import (
	"../object"
	"fmt"
	//"sort"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}

		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Integer{Value: int64(len(arg.Elements))}
		case *object.String:
			return &object.Integer{Value: int64(len(arg.Value))}
		default:
			return newError("argument to `len` not supported, got %s",
				args[0].Type())
		}
	},
	},
	"consoleOut": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return &object.String{Value: ""}
		},
	},
	//String Library
	"toString": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Integer:
				return &object.String{Value: strconv.FormatInt((arg.Value), 10)}
			default:
				return newError("arguments to must be INTEGER, got %s",
					args[0].Type())
			}

		},
	},
	//Random Library
	"random": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Integer:
				switch arg1 := args[1].(type) {
				case *object.Integer:
					rand.Seed(time.Now().UnixNano())
					min := arg.Value
					max := arg1.Value
					return &object.Integer{Value: int64(rand.Int63n(max-min+1) + min)}
				default:
					return newError("Second argument must be an INTEGER, got %s",
						args[1].Type())
				}
			default:
				return newError("First argument must be an INTEGER, got %s",
					args[0].Type())
			}
		},
	},
	//Prints first element of array
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	//Prints last element of array
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	//Prints all elements except the first element of array
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	//Push a value to the end of the array
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	//Sort an array
	"sort": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `sort` must be ARRAY, got %s",
					args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("argument to `sort` must be INTEGER, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length, length)
				copy(newElements, arr.Elements[0:length])
				//sort.Sort(sort.StringSlice(newElements[0:length]))
				//fmt.Println(object.Array{newElements})
				//WORK IN PROGRESS
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	//Math library
	"add": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return newError("wrong number of arguments. got=%d, want=1 or more",
					len(args))
			}
			for i := 0; i < len(args); i++ {
				if args[i].Type() != object.INTEGER_OBJ {
					return newError("argument to `add` must be INTEGER, got %s",
						args[0].Type())
				}
			}
			addition := int64(0)
			for i := 0; i < len(args); i++ {
				switch arg := args[i].(type) {
				case *object.Integer:
					addition = addition + arg.Value
				}
			}
			return &object.Integer{Value: int64(addition)}
		},
	},
	"help": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("wrong number of arguments. got=%d, want=none",
					len(args))
			}
			return &object.String{Value: ("For additional help please access the gitHub repository at https://github.com/ryanrasi/mocha-interpreter")}
		},
	},
	"split": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			} else {
				if args[0].Type() != object.STRING_OBJ {
					return newError("argument to `split` must be STRING, got %s",
						args[0].Type())
				} else if args[1].Type() != object.STRING_OBJ {
					return newError("argument to `split` must be STRING, got %s",
						args[1].Type())
				} else {
					var stringToSplit = ""
					var splitter = ""
					switch arg := args[0].(type) {
					case *object.String:
						stringToSplit = arg.Value
						print(stringToSplit)
						//print(arg.Value)
					default:
						return newError("arguments to must be STRING, got %s",
							args[0].Type())
					}
					switch arg := args[1].(type) {
					case *object.String:
						splitter = arg.Value
					default:
						return newError("arguments to must be STRING, got %s",
							args[0].Type())
					}
					var returnString = strings.Split(stringToSplit, splitter)
					print(returnString)
					//newElements := make([]object.Object, 2, 2)
					//newElements[0] = returnString[0]
					//newElements[1] = returnString[1]
					return NULL //&object.Array{Elements: newElements}
					//copy(newElements, returnString[0:len(returnString)])
					//sort.Sort(sort.StringSlice(newElements[0:length]))
					//fmt.Println(object.Array{newElements})
					//WORK IN PROGRESS
					//return &object.Array{Elements: newElements}
				}

			}
			return NULL
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("wrong number of arguments. got=%d, want=none",
					len(args))
			}
			fmt.Println("Program exited - Exit code 0")
			os.Exit(0)
			return NULL
		},
	},
}
