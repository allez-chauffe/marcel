import { Component,OnInit,ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-root',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'app works!';
  reload : number = 300;

  constructor(){}

  ngOnInit(){
    this.refreshPage(this.reload*1000);
  }

  refreshPage(reload){
    setTimeout(() => {
      console.log(reload);
      console.log("on reload la page :)");
      window.location.reload();
    },reload);
  }
}
