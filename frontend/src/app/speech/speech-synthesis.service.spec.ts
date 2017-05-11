import { TestBed, inject } from '@angular/core/testing';

import { SpeechSynthesisService } from './speech-synthesis.service';

describe('SpeechSynthesisService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [SpeechSynthesisService]
    });
  });

  it('should ...', inject([SpeechSynthesisService], (service: SpeechSynthesisService) => {
    expect(service).toBeTruthy();
  }));
});
