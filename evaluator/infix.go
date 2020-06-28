package evaluator

import (
	"regexp"

	"github.com/shric/monkey/object"
	"github.com/shric/monkey/token"
)

func evalInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	switch {
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalFloatIntegerInfixExpression(tok, left, right)
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalIntegerFloatInfixExpression(tok, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(tok, left, right)
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(tok, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(tok, left, right)
	case tok.Type == token.EQ:
		return nativeBoolToBooleanObject(left == right)
	case tok.Type == token.NOT_EQ:
		return nativeBoolToBooleanObject(left != right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(tok, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), tok.Type, right.Type())
	default:
		return newError("unknown token: %s %s %s",
			left.Type(), tok.Literal, right.Type())
	}
}

func evalBooleanInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value
	switch tok.Type {
	case token.AND:
		return &object.Boolean{Value: leftVal && rightVal}
	case token.OR:
		return &object.Boolean{Value: leftVal || rightVal}
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), tok.Type, right.Type())
	}
}

func evalFloatInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value
	switch tok.Type {
	case token.SLASH:
		if rightVal == 0 {
			return newError("Integer division by zero: %f/0.0", leftVal)
		}
		return &object.Float{Value: leftVal / rightVal}
	case token.PLUS:
		return &object.Float{Value: leftVal + rightVal}
	case token.MINUS:
		return &object.Float{Value: leftVal - rightVal}
	case token.ASTERISK:
		return &object.Float{Value: leftVal * rightVal}
	case token.LT:
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case token.GT:
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case token.EQ:
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), tok.Type, right.Type())
	}
}

func evalIntegerInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch tok.Type {
	case token.PLUS:
		return &object.Integer{Value: leftVal + rightVal}
	case token.MINUS:
		return &object.Integer{Value: leftVal - rightVal}
	case token.ASTERISK:
		return &object.Integer{Value: leftVal * rightVal}
	case token.SLASH:
		if rightVal == 0 {
			return newError("Integer division by zero: %d/0", leftVal)
		}
		return &object.Integer{Value: leftVal / rightVal}
	case token.LT:
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case token.GT:
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case token.EQ:
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), tok.Type, right.Type())
	}
}

func evalIntegerFloatInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Float).Value

	switch tok.Type {
	case token.PLUS:
		return &object.Float{Value: float64(leftVal) + rightVal}
	case token.MINUS:
		return &object.Float{Value: float64(leftVal) - rightVal}
	case token.ASTERISK:
		return &object.Float{Value: float64(leftVal) * rightVal}
	case token.SLASH:
		if rightVal == 0 {
			return newError("Integer division by zero: %d/0", leftVal)
		}
		return &object.Float{Value: float64(leftVal) / rightVal}
	case token.LT:
		return nativeBoolToBooleanObject(float64(leftVal) < rightVal)
	case token.GT:
		return nativeBoolToBooleanObject(float64(leftVal) > rightVal)
	case token.EQ:
		return nativeBoolToBooleanObject(float64(leftVal) == rightVal)
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(float64(leftVal) != rightVal)
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), tok.Type, right.Type())
	}
}

func evalFloatIntegerInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Integer).Value

	switch tok.Type {
	case token.PLUS:
		return &object.Float{Value: leftVal + float64(rightVal)}
	case token.MINUS:
		return &object.Float{Value: leftVal - float64(rightVal)}
	case token.ASTERISK:
		return &object.Float{Value: leftVal * float64(rightVal)}
	case token.SLASH:
		if rightVal == 0 {
			return newError("Integer division by zero: %f/0", leftVal)
		}
		return &object.Float{Value: leftVal / float64(rightVal)}
	case token.LT:
		return nativeBoolToBooleanObject(leftVal < float64(rightVal))
	case token.GT:
		return nativeBoolToBooleanObject(leftVal > float64(rightVal))
	case token.EQ:
		return nativeBoolToBooleanObject(leftVal == float64(rightVal))
	case token.NOT_EQ:
		return nativeBoolToBooleanObject(leftVal != float64(rightVal))
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), tok.Type, right.Type())
	}
}

func evalStringInfixExpression(
	tok token.Token,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch tok.Type {
	case token.PLUS:
		return &object.String{Value: leftVal + rightVal}
	case token.EQ:
		return &object.Boolean{Value: leftVal == rightVal}
	case token.NOT_EQ:
		return &object.Boolean{Value: leftVal != rightVal}
	case token.REGEX:
		re, err := regexp.Compile(rightVal)
		if err != nil {
			return newError("%v", err)
		}
		return nativeBoolToBooleanObject(re.MatchString(leftVal))
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), tok.Type, right.Type())
	}
}
