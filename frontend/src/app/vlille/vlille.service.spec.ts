/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { VlilleService } from './vlille.service';

describe('Service: Vlille', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [VlilleService]
    });
  });

  it('should ...', inject([VlilleService], (service: VlilleService) => {
    expect(service).toBeTruthy();
  }));
});
