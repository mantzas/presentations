"use strict";

import { APIGatewayEvent, Callback, Context, ProxyResult } from "aws-lambda";
import { OrderProcessor } from "./OrderProcessor";
import { Repository } from "./Repository";

interface IEventJournal {
  timestamp: string;
  title: string;
  description: string;
  source: string;
}

const processor = new OrderProcessor(new Repository());

export async function retrieve(event: APIGatewayEvent, context: Context, callback: Callback) {

  try {

    const orders = await processor.retrieve();

    console.log(`got ${orders.length} orders`);

    callback(null, createSuccessResponse(orders));
  } catch (error) {
    console.log(error);
    callback(null, createInternalServerErrorResponse("failed to retrieve orders!"));
  }
}

function createSuccessResponse(data?): ProxyResult {
  return {
    body: data == null ? null : JSON.stringify(data),
    headers: getDefaultHeader(),
    statusCode: 200,
  };
}

function createInternalServerErrorResponse(message: string): ProxyResult {

  return {
    body: JSON.stringify({
      error: message,
    }),
    headers: getDefaultHeader(),
    statusCode: 500,
  };
}

function getDefaultHeader() {
  return {
    ["Content-Type"]: "application/json",
  };
}
