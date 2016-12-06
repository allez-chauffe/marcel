// The file contents for the current environment will overwrite these during build.
// The build system defaults to the dev environment which uses `environment.ts`, but if you do
// `ng build --env=prod` then `environment.prod.ts` will be used instead.
// The list of which env maps to which file can be found in `angular-cli.json`.

export const environment = {
  production: false,

  weatherUrl: 'http://10.0.10.63:8090/api/v1/weather/forecast/5',
  weatherKey : 'YOUR_FREE_OPENWEATHER_API_KEY'
};
