export interface IOrder {
    id: number;
    items: IOrderItem[];
    price: number;
}

export interface IOrderItem {
    id: number;
    product: string;
    quantity: number;
    price: number;
}

export class Repository {
    public getOrders(): Promise<IOrder[]> {

        const orders = new Array<IOrder>();

        orders.push({
            id: 1,
            items: [{
                id: 1,
                price: 2.99,
                product: "Product 1",
                quantity: 2,
            }, {
                id: 2,
                price: 1.99,
                product: "Product 2",
                quantity: 1,
            }],
            price: 7.97,
        });

        orders.push({
            id: 2,
            items: [{
                id: 1,
                price: 1.99,
                product: "Product 2",
                quantity: 2,
            }],
            price: 1.99,
        });

        return Promise.resolve(orders);
    }
}
