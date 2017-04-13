import { Component, OnInit, ViewEncapsulation, OnDestroy } from '@angular/core';
import { SpeechRecognitionService } from './speech-recognition.service';

@Component({
  selector: 'speech',
  templateUrl: './speech.component.html',
  styleUrls: ['./speech.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class SpeechComponent implements OnInit, OnDestroy {
  speechData: String;

  constructor(private speechRecognitionService: SpeechRecognitionService) {
    this.speechData = "Hello";
  }

  ngOnInit() {
    this.startRecognition();
  }

  startRecognition() {
    this.speechRecognitionService.record()
    .subscribe(
      (value) => {
        this.speechData = value;
        console.log(value);
      },
      (err) => {
        console.log(err);
        if (err.error == "no-speech") {
          console.log("--restarting service--");
          this.startRecognition();
        }
      },
      //completion
      () => {
        console.log("--complete--");
        this.startRecognition();
      });
  }

  ngOnDestroy() {
    this.speechRecognitionService.DestroySpeechObject();
  }

}
