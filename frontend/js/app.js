'use strict';

/*-----------------
 APP
 -----------------*/

let app = null;

function startVue() {
	app = new Vue({
		el: '#componentsList',

		data: {
			pluginsData: {}
		},

		created : function () {
			asyncComponentLoader.setPropsValues(this);
		},
		computed: {}
	});
}
