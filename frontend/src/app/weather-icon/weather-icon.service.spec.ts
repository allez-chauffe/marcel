/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { WeatherIconService } from './weather-icon.service';

describe('Service: WeatherIcon', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [WeatherIconService]
    });
  });

  it('should ...', inject([WeatherIconService], (service: WeatherIconService) => {
    expect(service).toBeTruthy();
  }));
});
