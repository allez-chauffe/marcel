import { Component, OnInit, ViewEncapsulation, OnDestroy } from '@angular/core';
import { SpeechRecognitionService } from './speech-recognition.service';
import { ApiAiClient } from 'api-ai-javascript/ApiAiClient';
import { ApiService } from './../api/api.service';
import { YoutubeService } from './../youtube/youtube.service';

@Component({
  selector: 'speech',
  templateUrl: './speech.component.html',
  styleUrls: ['./speech.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class SpeechComponent implements OnInit, OnDestroy {

  speechData: String;
  apiAiClient: ApiAiClient;

  constructor(private speechRecognitionService: SpeechRecognitionService, private youtubeService: YoutubeService, private apiService: ApiService) {
    this.speechData = "Hello";
    this.apiAiClient = new ApiAiClient({accessToken: apiService.getApi('apiai').token });
  }

  ngOnInit() {
    this.startRecognition();
    this.sendRequest('Montre moi des vidéos de chat');
  }

  startRecognition() {
    this.speechRecognitionService.record()
    .subscribe(
      (value) => {
        this.speechData = value;
        this.sendRequest(value);
      },
      (err) => {
        if (err.error == "no-speech") {
          console.log("--restarting service--");
        } else {
          console.log(err);
        }
        this.startRecognition();
      },
      //completion
      () => {
        console.log("--complete--");
        this.startRecognition();
      });
  }

  sendRequest(request) {
    if(request.length !== 0){
      this.apiAiClient
        .textRequest(request)
        .then((response) => {
          let test: any = response;
          this.handleRequest(test.result.parameters, test.result.fulfillment);
        })
        .catch((error) => {
          console.log(error);
        })
    }
  }

  handleRequest(parameters: any, fulfillment: any) {
    if(parameters.video !== undefined){
      console.log(fulfillment.speech);
      this.youtubeService.startSearch(parameters.video);
    }
  }

  ngOnDestroy() {
    this.speechRecognitionService.DestroySpeechObject();
  }

}
