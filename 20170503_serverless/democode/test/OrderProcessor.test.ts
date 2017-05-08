const expect = require("chai").expect;
import { OrderProcessor } from "../src/OrderProcessor";
import { IOrder, Repository } from "../src/Repository";

describe("OrderProcessor", () => {
    describe("retrieve", () => {
        it("succeeds", async () => {

            const repository = new Repository();
            repository.getOrders = () => {
                const orders: IOrder[] = [{
                    id: 1,
                    items: [{
                        id: 1,
                        price: 1.1,
                        product: "XXX",
                        quantity: 1,
                    }],
                    price: 1.1,
                }];

                return Promise.resolve(orders);
            };
            const orders = await new OrderProcessor(repository).retrieve();

            expect(orders.length).to.equal(1);
            expect(orders[0].id).to.equal(1);
            expect(orders[0].price).to.equal(1.1);
        });
    });
});
