package handler

import (
	"cart/common"
	"cart/domain/model"
	"cart/domain/service"
	cart "cart/proto"
	"context"

)

type Cart struct{
	CateDatService service.ICartDataService
}


func (c *Cart)AddCart(ctx context.Context, req *cart.CartInfo, res *cart.ResponseAdd) (err error){
	cart :=&model.Cart{}
	if err = common.SwapTo(req,cart);err != nil{
		return err
	}
	if res.CartId,err = c.CateDatService.AddCart(cart);err != nil{
		return err
	}
	res.Msg = "新增商品成功"
	return nil
}
func (c *Cart)CleanCart(ctx context.Context, req * cart.Clean, res *cart.Response) error{
	err := c.CateDatService.CleanCart(req.UserId)
	if err != nil {
		return err
	}
	res.Msg = "删除成功"
	return nil
}
func (c *Cart)Incr(ctx context.Context, req *cart.Item, res *cart.Response) error{
	err  := c.CateDatService.IncrNum(req.Id,req.ChangeNum)
	if err != nil {
		return err
	}
	res.Msg = "增加成功"
	return nil
}
func (c *Cart)Decr(ctx context.Context, req *cart.Item, res *cart.Response) error{
	err  := c.CateDatService.IncrNum(req.Id,req.ChangeNum)
	if err != nil {
		return err
	}
	res.Msg = "减库成功"
	return nil
}
func (c *Cart)DeleteItemByID(ctx context.Context, req *cart.CartID, res *cart.Response) error{
	err := c.CateDatService.DeleteCart(req.CartId)
	if err != nil {
		return err
	}
	res.Msg = "删除成功"
	return nil
}
func (c *Cart)GetAll(ctx context.Context, req *cart.CartFindAll, res *cart.CartAll) error{
	data,err := c.CateDatService.FindAllCart(req.UserId)
	if err != nil {
		return err
	}
	err = common.SwapTo(data,res.CartInfo)
	if err != nil {
		return err
	}
	return nil
}