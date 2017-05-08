"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var OrderProcessor = (function () {
    function OrderProcessor(repository) {
        this.repository = repository;
    }
    OrderProcessor.prototype.retrieve = function () {
        return this.repository.getOrders();
    };
    return OrderProcessor;
}());
exports.OrderProcessor = OrderProcessor;
//# sourceMappingURL=OrderProcessor.js.map