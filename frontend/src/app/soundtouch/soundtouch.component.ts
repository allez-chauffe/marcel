import { Component, OnInit,ViewEncapsulation } from '@angular/core';
import { SoundtouchService} from './soundtouch.service'

@Component({
  encapsulation: ViewEncapsulation.None,
  selector: 'soundtouch',
  templateUrl: './soundtouch.component.html',
  styleUrls: ['./soundtouch.component.scss']
})
export class SoundtouchComponent implements OnInit {

  sound : any;
  private timer: number = 1000 * 60 * 1;

  constructor(public soundtouchService : SoundtouchService) { }

  ngOnInit() {
    setInterval(() => {
      this.getNowPlaying();
    }, this.timer);
    this.getNowPlaying();
  }

  getNowPlaying(){
    this.soundtouchService.getNowPlaying().subscribe( data => {
      console.log(data);
      this.sound = data;
    });
  }

}
