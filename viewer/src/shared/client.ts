import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { WorkloadService, type CreateRequest } from '../generated/workload/v1/workload_pb.ts';

export class Client {
  private tp: any

  constructor() {
    this.tp = createConnectTransport({
      baseUrl: "http://localhost:8080",
      useBinaryFormat: true,
    });
  }

  Workload(): any {
    return new Workload(this.tp);
  }
}

export class Workload {
  private client

  constructor(tp: any) {
    this.client = createClient(WorkloadService, tp);
    console.log(this.client)
  }

  async Get(id: string) {
    return this.client.get({ id: id });
  }

  async List() {
    return this.client.list({});
  }

  async Create(request: CreateRequest) {
    return this.client.create(request);
  }

  async Delete(id: string) {
    return this.client.delete({ id: id });
  }
}
