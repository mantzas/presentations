import { IOrder, Repository } from "./Repository";

export class OrderProcessor {

    private readonly repository: Repository;

    constructor(repository: Repository) {
        this.repository = repository;
    }

    public retrieve(): Promise<IOrder[]> {

        return this.repository.getOrders();
    }
}
