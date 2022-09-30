export namespace kafka {
	
	export class PartitionOffset {
	    partition: number;
	    committedOffset: number;
	    metadata: string;
	
	    static createFrom(source: any = {}) {
	        return new PartitionOffset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partition = source["partition"];
	        this.committedOffset = source["committedOffset"];
	        this.metadata = source["metadata"];
	    }
	}
	export class ConsumerOffset {
	    consumer: string;
	    offsets: PartitionOffset[];
	
	    static createFrom(source: any = {}) {
	        return new ConsumerOffset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.consumer = source["consumer"];
	        this.offsets = this.convertValues(source["offsets"], PartitionOffset);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ConfigEntry {
	    name: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new ConfigEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	    }
	}
	export class ReplicaAssignment {
	    partition: number;
	    replicas: number[];
	
	    static createFrom(source: any = {}) {
	        return new ReplicaAssignment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partition = source["partition"];
	        this.replicas = source["replicas"];
	    }
	}
	export class TopicConfig {
	    topic: string;
	    numPartitions: number;
	    replicationFactor: number;
	    replicaAssignments: ReplicaAssignment[];
	    configEntries: ConfigEntry[];
	
	    static createFrom(source: any = {}) {
	        return new TopicConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topic = source["topic"];
	        this.numPartitions = source["numPartitions"];
	        this.replicationFactor = source["replicationFactor"];
	        this.replicaAssignments = this.convertValues(source["replicaAssignments"], ReplicaAssignment);
	        this.configEntries = this.convertValues(source["configEntries"], ConfigEntry);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Broker {
	    host: string;
	    port: number;
	    id: number;
	    rack: string;
	
	    static createFrom(source: any = {}) {
	        return new Broker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.host = source["host"];
	        this.port = source["port"];
	        this.id = source["id"];
	        this.rack = source["rack"];
	    }
	}
	export class Partition {
	    topic: string;
	    id: number;
	    // Go type: Broker
	    leader: any;
	    replicas: Broker[];
	
	    static createFrom(source: any = {}) {
	        return new Partition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topic = source["topic"];
	        this.id = source["id"];
	        this.leader = this.convertValues(source["leader"], null);
	        this.replicas = this.convertValues(source["replicas"], Broker);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Topic {
	    name: string;
	    internal: boolean;
	    partitions: Partition[];
	
	    static createFrom(source: any = {}) {
	        return new Topic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.internal = source["internal"];
	        this.partitions = this.convertValues(source["partitions"], Partition);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace application {
	
	export class Profile {
	    name: string;
	    bootstrapServers: string[];
	
	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.bootstrapServers = source["bootstrapServers"];
	    }
	}
	export class Config {
	    version: string;
	    profiles: Profile[];
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.profiles = this.convertValues(source["profiles"], Profile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

