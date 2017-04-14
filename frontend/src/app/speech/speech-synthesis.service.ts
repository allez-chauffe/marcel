import { Injectable } from '@angular/core';

declare var Espeak: any;
declare var PushAudioNode: any;

@Injectable()
export class SpeechSynthesisService {
  ctx: AudioContext;
  convolver: any;
  espeak: any;
  static pusher: any;
  now: any;

  constructor() {
    var win : any = window;
    this.ctx = new (win.AudioContext || win.webkitAudioContext)();

    this.espeak = new Espeak('espeak/js/espeak.worker.js', () => {
    })
    // this.convolver = this.ctx.createConvolver();
    // this.convolver.connect(this.ctx.destination)
  }

  stop() {
    if(SpeechSynthesisService.pusher) {
      SpeechSynthesisService.pusher.disconnect();
      SpeechSynthesisService.pusher = null;
    }
  }

  speak(message: String) {
    this.stop();
    var samples_queues : Array<any> = [];
    this.espeak.set_rate(150);
    this.espeak.set_pitch(200);
    this.espeak.setVoice.apply(this.espeak, ['french', 'fr']);
    this.now = Date.now();
    SpeechSynthesisService.pusher = new PushAudioNode(this.ctx);
    SpeechSynthesisService.pusher.connect(this.ctx.destination);
    this.espeak.synth(message, this.handleSynth);
  }

  handleSynth(samples, events) {
    if (!samples) {
      SpeechSynthesisService.pusher.close();
      return;
    }
    SpeechSynthesisService.pusher.push(new Float32Array(samples));
    //if (events.length)
    //  console.log(events.map((e) => e.type));
    if (this.now) {
      console.log("latency:", Date.now() - this.now);
      this.now = 0;
    }
  }

}
