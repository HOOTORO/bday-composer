export namespace main {
	
	export class Contact {
	    name: string;
	    birthday: string;
	
	    static createFrom(source: any = {}) {
	        return new Contact(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.birthday = source["birthday"];
	    }
	}

}

