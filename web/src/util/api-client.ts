import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { WorkloadService, CreateRequest } from "@/generated/workload/v1/workload_pb";
import { VersionService, CreateRequest as VCreate } from "@/generated/version/v1/version_pb";

export default class APIClient {
  private tp: any

  constructor() {
    this.tp = createConnectTransport({
      baseUrl: "http://localhost:8080",
      useBinaryFormat: true,
    });
  }

  Workload(): Workload {
    return new Workload(this.tp);
  }

  Version(): Version {
    return new Version(this.tp);
  }
}

export class Workload {
  private client

  constructor(tp: any) {
    this.client = createClient(WorkloadService, tp);
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

export class Version {
  private client

  constructor(tp: any) {
    this.client = createClient(VersionService, tp);
  }

  async Get(id: string) {
    return this.client.get({ id: id });
  }

  async List() {
    return this.client.list({});
  }

  async Create(request: VCreate) {
    return this.client.create(request);
  }

  async Delete(id: string) {
    return this.client.delete({ id: id });
  }
}
