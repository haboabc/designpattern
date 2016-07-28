/* 超级工厂factoryproducer, 对外提供的服务StartService() 和 GetFactory() */
/* 多有工厂(超级)都会提供2中有用的东西，服务和机器(工厂) */
package factoryproducer

import "fmt"

/* 空接口是被任何类型都默认已经实现的 */
type Fcproducer interface {
	/*因为要用的功能可能是Draw()或者Fill()
	 *只能起一个笼统的个功能名字，包含2者
	 */
	 /* 用户将只能看到这个功能 */
	 StartService(param string)
	 /* 上边这个功能会去调用底层的Draw(), Fill() */
}

/*
type ShapeFactory struct {
}

type ColorFactory struct {
}
*/

/* 根据不同选择返回不同类型 , 用同一个藉口接收 */
func GetFactory(choice string)  Fcproducer{
	switch(choice) {
		case "color":
			return ColorFactory{}
		case "shape":
			return ShapeFactory{}
		default:
			fmt.Println("unknown choice")
			return nil
	}
}

func StartService(choice, param string) {
	var factory Fcproducer

	factory = GetFactory(choice)

	if factory != nil {
		/* 这个接口目前只有这一种方法 */
		factory.StartService(param)
	} else {
		//err
		//already done in the GetFactory()
	}
}
