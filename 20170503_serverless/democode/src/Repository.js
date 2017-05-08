"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var Repository = (function () {
    function Repository() {
    }
    Repository.prototype.getOrders = function () {
        var orders = new Array();
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
    };
    return Repository;
}());
exports.Repository = Repository;
//# sourceMappingURL=Repository.js.map