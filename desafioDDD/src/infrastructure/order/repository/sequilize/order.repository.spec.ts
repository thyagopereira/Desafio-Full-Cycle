import { Sequelize } from "sequelize-typescript";
import Order from "../../../../domain/checkout/entity/order";
import OrderItem from "../../../../domain/checkout/entity/order_item";
import Customer from "../../../../domain/customer/entity/customer";
import Address from "../../../../domain/customer/value-object/address";
import Product from "../../../../domain/product/entity/product";
import CustomerModel from "../../../customer/repository/sequelize/customer.model";
import CustomerRepository from "../../../customer/repository/sequelize/customer.repository";
import ProductModel from "../../../product/repository/sequelize/product.model";
import ProductRepository from "../../../product/repository/sequelize/product.repository";
import OrderItemModel from "./order-item.model";
import OrderModel from "./order.model";
import OrderRepository from "./order.repository";

describe("Order repository test", () => {
  let sequelize: Sequelize;

  beforeEach(async () => {
    sequelize = new Sequelize({
      dialect: "sqlite",
      storage: ":memory:",
      logging: false,
      sync: { force: true },
    });

    await sequelize.addModels([
      CustomerModel,
      OrderModel,
      OrderItemModel,
      ProductModel,
    ]);
    await sequelize.sync();
  });

  afterEach(async () => {
    await sequelize.close();
  });

  it("should create a new order", async () => {
    const customerRepository = new CustomerRepository();
    const customer = new Customer("123", "Customer 1");
    const address = new Address("Street 1", 1, "Zipcode 1", "City 1");
    customer.changeAddress(address);
    await customerRepository.create(customer);

    const productRepository = new ProductRepository();
    const product = new Product("123", "Product 1", 10);
    await productRepository.create(product);

    const orderItem = new OrderItem(
      "1",
      product.name,
      product.price,
      product.id,
      2
    );

    const order = new Order("123", "123", [orderItem]);

    const orderRepository = new OrderRepository();
    await orderRepository.create(order);

    const orderModel = await OrderModel.findOne({
      where: { id: order.id },
      include: ["items"],
    });

    expect(orderModel.toJSON()).toStrictEqual({
      id: "123",
      customer_id: "123",
      total: order.total(),
      items: [
        {
          id: orderItem.id,
          name: orderItem.name,
          price: orderItem.price,
          quantity: orderItem.quantity,
          order_id: "123",
          product_id: "123",
        },
      ],
    });
  });

  it("Should find a existing order", async () => {
    const customerRepository = new CustomerRepository();
    const customer = new Customer("123", "Customer 1");
    const address = new Address("Street 1", 1, "Zipcode 1", "City 1");
    customer.changeAddress(address);
    await customerRepository.create(customer);

    const productRepository = new ProductRepository();
    const product = new Product("123", "Product 1", 10);
    await productRepository.create(product);

    const orderItem = new OrderItem(
      "1",
      product.name,
      product.price,
      product.id,
      2
    );

    const order = new Order("123", "123", [orderItem]);

    const orderRepository = new OrderRepository();
    await orderRepository.create(order);

    const result = await orderRepository.find(order.id);

    expect(result).toStrictEqual(order);
  });

  it("Should find all Existing Orders", async () => {
    const customerRepository = new CustomerRepository();
    const customer = new Customer("123", "Customer 1");
    const address = new Address("Street 1", 1, "Zipcode 1", "City 1");
    customer.changeAddress(address);
    await customerRepository.create(customer);

    const customerTwo = new Customer("345", "Customer 2");
    const addresTwo = new Address("Street 2", 2, "Zipcode 2", "City 2");
    customerTwo.changeAddress(addresTwo);
    await customerRepository.create(customerTwo);

    const productRepository = new ProductRepository();
    const product = new Product("123", "Product 1", 10);
    const producttwo = new Product("345", "Product 2 ", 20);
    await productRepository.create(product);
    await productRepository.create(producttwo);

    const orderItem = new OrderItem(
      "1",
      product.name,
      product.price,
      product.id,
      2
    );

    const orderItemTwo =  new OrderItem("2", producttwo.name , producttwo.price, producttwo.id, 3);

    const order = new Order("123", "123", [orderItem]);

    const orderRepository = new OrderRepository();
    await orderRepository.create(order);

    const orderTwo = new Order("345", "345", [orderItemTwo]);
    await orderRepository.create(orderTwo);

    const foundOrders = await orderRepository.findAll();
    const orders = [order, orderTwo]

    expect(orders).toEqual(foundOrders)
  });


  it("Should update a existing order", async () => {

    // Set up customer 
    const customerRepository = new CustomerRepository();
    const customer = new Customer("123", "Customer 1");
    const address = new Address("Street 1", 1, "Zipcode 1", "City 1");
    customer.changeAddress(address);
    await customerRepository.create(customer);
    
    // Set up product
    const productRepository = new ProductRepository();
    const product = new Product("123", "Product 1", 10);
    await productRepository.create(product);

    // Set up OrderItem
    const orderItem = new OrderItem(
      "1",
      product.name,
      product.price,
      product.id,
      2
    );

    //Finally set up Order. 
    const order = new Order("123", "123", [orderItem]);
    const orderRepository = new OrderRepository();
    await orderRepository.create(order);


    // Creates new orderItem
    const updatedCustomer = new Customer("456", "updated");
    const updatedProduct  = new Product("456", "Product 2", 20);
    const updatedProductTwo =  new Product("789", "Product 3", 30);
    const updatedItems = [new OrderItem("2", updatedProduct.name, updatedProduct.price, 
      updatedProduct.id, 4), new OrderItem("3", updatedProductTwo.name, 
      updatedProductTwo.price, updatedProductTwo.id, 5)]; 

    order.changeCustomer(updatedCustomer);
    order.changeItems(updatedItems);

    orderRepository.update(order);
    const updatedOrder = await orderRepository.find(order.id);

    expect(updatedOrder).toStrictEqual(order);
        
  }); 
});
