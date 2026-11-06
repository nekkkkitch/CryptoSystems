export namespace app {
	
	export class KeysToRet {
	    pu: number;
	    pr: number;
	
	    static createFrom(source: any = {}) {
	        return new KeysToRet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pu = source["pu"];
	        this.pr = source["pr"];
	    }
	}

}

