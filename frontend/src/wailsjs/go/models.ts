export namespace main {
	
	export class ColumnInfo {
	    name: string;
	    type: string;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new ColumnInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.size = source["size"];
	    }
	}
	export class TablelInfo {
	    table_name: string;
	
	    static createFrom(source: any = {}) {
	        return new TablelInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.table_name = source["table_name"];
	    }
	}

}

export namespace model {
	
	export class BuildProject {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    name: string;
	    project_path: string;
	    build_type: number;
	    version: string;
	    build_version: number;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new BuildProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.name = source["name"];
	        this.project_path = source["project_path"];
	        this.build_type = source["build_type"];
	        this.version = source["version"];
	        this.build_version = source["build_version"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class CodeTemplate {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    fileName: string;
	    schemeCode: string;
	    schemeParentCode: string;
	    mark: string;
	    nodeType: number;
	    fileType: number;
	    icon: string;
	
	    static createFrom(source: any = {}) {
	        return new CodeTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.fileName = source["fileName"];
	        this.schemeCode = source["schemeCode"];
	        this.schemeParentCode = source["schemeParentCode"];
	        this.mark = source["mark"];
	        this.nodeType = source["nodeType"];
	        this.fileType = source["fileType"];
	        this.icon = source["icon"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class DataModel {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    makeType: number;
	    tableName: string;
	    project: string;
	    schemeId: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new DataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.makeType = source["makeType"];
	        this.tableName = source["tableName"];
	        this.project = source["project"];
	        this.schemeId = source["schemeId"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class DataSource {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    host: string;
	    port: number;
	    database: string;
	    user_name: string;
	    password: string;
	    db_type: string;
	    schema: string;
	    path: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new DataSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user_name = source["user_name"];
	        this.password = source["password"];
	        this.db_type = source["db_type"];
	        this.schema = source["schema"];
	        this.path = source["path"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ModelInfo {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    dataModelId: string;
	    title: string;
	    name: string;
	    columnType: string;
	    size: number;
	    isPk: number;
	    extension: string[];
	    controlType: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new ModelInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.dataModelId = source["dataModelId"];
	        this.title = source["title"];
	        this.name = source["name"];
	        this.columnType = source["columnType"];
	        this.size = source["size"];
	        this.isPk = source["isPk"];
	        this.extension = source["extension"];
	        this.controlType = source["controlType"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Scheme {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    code: string;
	    title: string;
	    name: string;
	    level: number;
	    parentCode: string;
	    icon: string;
	    nodeType: number;
	    nodes: string[];
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new Scheme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.code = source["code"];
	        this.title = source["title"];
	        this.name = source["name"];
	        this.level = source["level"];
	        this.parentCode = source["parentCode"];
	        this.icon = source["icon"];
	        this.nodeType = source["nodeType"];
	        this.nodes = source["nodes"];
	        this.path = source["path"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Variable {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    pro_id: string;
	    key: string;
	    value: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new Variable(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.pro_id = source["pro_id"];
	        this.key = source["key"];
	        this.value = source["value"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class VariableBase {
	    pro_id: string;
	    key: string;
	    value: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new VariableBase(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pro_id = source["pro_id"];
	        this.key = source["key"];
	        this.value = source["value"];
	        this.mark = source["mark"];
	    }
	}

}

export namespace repository {
	
	export class CreateBuildProject {
	    title: string;
	    name: string;
	    project_path: string;
	    build_type: number;
	    version: string;
	    build_version: number;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateBuildProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.name = source["name"];
	        this.project_path = source["project_path"];
	        this.build_type = source["build_type"];
	        this.version = source["version"];
	        this.build_version = source["build_version"];
	        this.mark = source["mark"];
	    }
	}
	export class CreateDataModel {
	    title: string;
	    makeType: number;
	    tableName: string;
	    project: string;
	    schemeId: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateDataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.makeType = source["makeType"];
	        this.tableName = source["tableName"];
	        this.project = source["project"];
	        this.schemeId = source["schemeId"];
	        this.mark = source["mark"];
	    }
	}
	export class CreateDataSource {
	    title: string;
	    host: string;
	    port: number;
	    database: string;
	    user_name: string;
	    password: string;
	    db_type: string;
	    schema: string;
	    path: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateDataSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user_name = source["user_name"];
	        this.password = source["password"];
	        this.db_type = source["db_type"];
	        this.schema = source["schema"];
	        this.path = source["path"];
	        this.mark = source["mark"];
	    }
	}
	export class CreateTemplate {
	    title: string;
	    fileName: string;
	    schemeCode: string;
	    schemeParentCode: string;
	    mark: string;
	    nodeType: number;
	    fileType: number;
	    icon: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.fileName = source["fileName"];
	        this.schemeCode = source["schemeCode"];
	        this.schemeParentCode = source["schemeParentCode"];
	        this.mark = source["mark"];
	        this.nodeType = source["nodeType"];
	        this.fileType = source["fileType"];
	        this.icon = source["icon"];
	    }
	}
	export class ModifyBuildProject {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    name: string;
	    project_path: string;
	    build_type: number;
	    version: string;
	    build_version: number;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new ModifyBuildProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.name = source["name"];
	        this.project_path = source["project_path"];
	        this.build_type = source["build_type"];
	        this.version = source["version"];
	        this.build_version = source["build_version"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ModifyTemplate {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    fileName: string;
	    schemeCode: string;
	    schemeParentCode: string;
	    mark: string;
	    nodeType: number;
	    fileType: number;
	    icon: string;
	
	    static createFrom(source: any = {}) {
	        return new ModifyTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.fileName = source["fileName"];
	        this.schemeCode = source["schemeCode"];
	        this.schemeParentCode = source["schemeParentCode"];
	        this.mark = source["mark"];
	        this.nodeType = source["nodeType"];
	        this.fileType = source["fileType"];
	        this.icon = source["icon"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class QryBuildProject {
	    title: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new QryBuildProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.name = source["name"];
	    }
	}
	export class QryTemplate {
	    condition: string;
	    parentCode: string;
	    ids: string[];
	
	    static createFrom(source: any = {}) {
	        return new QryTemplate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.condition = source["condition"];
	        this.parentCode = source["parentCode"];
	        this.ids = source["ids"];
	    }
	}
	export class RequestModelInfoSave {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    dataModelId: string;
	    title: string;
	    name: string;
	    columnType: string;
	    size: number;
	    isPk: number;
	    extension: string[];
	    controlType: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestModelInfoSave(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.dataModelId = source["dataModelId"];
	        this.title = source["title"];
	        this.name = source["name"];
	        this.columnType = source["columnType"];
	        this.size = source["size"];
	        this.isPk = source["isPk"];
	        this.extension = source["extension"];
	        this.controlType = source["controlType"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class RequestModifyDataModel {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    makeType: number;
	    tableName: string;
	    project: string;
	    schemeId: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new RequestModifyDataModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.makeType = source["makeType"];
	        this.tableName = source["tableName"];
	        this.project = source["project"];
	        this.schemeId = source["schemeId"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class SchemeTree {
	    id: string;
	    title: string;
	    icon: string;
	    type: number;
	    parentCode: string;
	    path: string;
	    children: SchemeTree[];
	
	    static createFrom(source: any = {}) {
	        return new SchemeTree(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.icon = source["icon"];
	        this.type = source["type"];
	        this.parentCode = source["parentCode"];
	        this.path = source["path"];
	        this.children = this.convertValues(source["children"], SchemeTree);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class UpdateDataSource {
	    id: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	    title: string;
	    host: string;
	    port: number;
	    database: string;
	    user_name: string;
	    password: string;
	    db_type: string;
	    schema: string;
	    path: string;
	    mark: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateDataSource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.title = source["title"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.database = source["database"];
	        this.user_name = source["user_name"];
	        this.password = source["password"];
	        this.db_type = source["db_type"];
	        this.schema = source["schema"];
	        this.path = source["path"];
	        this.mark = source["mark"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

