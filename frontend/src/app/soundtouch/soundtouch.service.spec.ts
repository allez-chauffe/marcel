/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { SoundtouchService } from './soundtouch.service';

describe('Service: Soundtouch', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [SoundtouchService]
    });
  });

  it('should ...', inject([SoundtouchService], (service: SoundtouchService) => {
    expect(service).toBeTruthy();
  }));
});
