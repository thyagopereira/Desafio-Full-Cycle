import Order from "../../../../domain/checkout/entity/order";
import OrderItem from "../../../../domain/checkout/entity/order_item";
import OrderRepositoryInterface from "../../../../domain/checkout/repository/order-repository.interface";
import OrderItemModel from "./order-item.model";
import orderModel from "./order.model";
import OrderModel from "./order.model";

export default class OrderRepository implements OrderRepositoryInterface{
  async create(entity: Order): Promise<void> {
    await OrderModel.create(
      {
        id: entity.id,
        customer_id: entity.customerId,
        total: entity.total(),
        items: entity.items.map((item) => ({
          id: item.id,
          name: item.name,
          price: item.price,
          product_id: item.productId,
          quantity: item.quantity,
        })),
      },
      {
        include: [{ model: OrderItemModel }],
      }
    );
  }
  
  async update(entity: Order): Promise<void> { 

  }

  async find(id: string): Promise<Order> {
    let orderModel;

    try{
      orderModel = await OrderModel.findOne({
        where: {
          id: id,
        },
        include: ["items"],
        rejectOnEmpty: true,
      });
    }catch (error) {
      throw new Error("Order  not found");
    }

    const itens = orderModel.items.map((orderModelItem) => {
      return new OrderItem(orderModelItem.id, orderModelItem.name, orderModelItem.price,
        orderModelItem.product_id, orderModelItem.quantity);
    });
    const order = new Order(id, orderModel.id, itens);
    return order
  }

  async findAll(): Promise<Order[]> {
    const orderModels = await OrderModel.findAll({include : "items"});
    
    return orderModels.map((orderModel) => {
      return new Order(orderModel.id, orderModel.customer_id, 
        orderModel.items.map((orderModelItem) => {

          return new OrderItem(orderModelItem.id, orderModelItem.name, orderModelItem.price,
            orderModelItem.product_id, orderModelItem.quantity); }))         
  });
}

}