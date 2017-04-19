import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { SpeechComponent } from './speech.component';
import { SpeechRecognitionService } from './speech-recognition.service';
import { SpeechSynthesisService } from './speech-synthesis.service';

@NgModule({
  imports:      [ CommonModule, FormsModule ],
  declarations: [ SpeechComponent ],
  exports:      [ SpeechComponent ],
  providers:    [ SpeechRecognitionService, SpeechSynthesisService ]
})
export class SpeechModule { }
