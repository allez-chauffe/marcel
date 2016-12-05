import { DateTimeModule } from './datetime/datetime.module';
import { GithubModule } from './github/github.module';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule,CUSTOM_ELEMENTS_SCHEMA} from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule,JsonpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { MessageComponent } from './message/message.component';
import { VlilleComponent } from './vlille/vlille.component';
import { VlilleService } from './vlille/vlille.service';
import { FirstletterupperPipe } from './pipes/firstletterupper.pipe';
import { ClockComponent } from './clock/clock.component';
import { HomeComponent } from './home/home.component';
import { SocialComponent } from './social/social.component';
import { CalendarComponent } from './calendar/calendar.component';
import { CalendarService } from './calendar/calendar.service';
import { WeatherComponent } from './weather/weather.component';
import { WeatherService } from './weather/weather.service';
import { WeatherIconComponent } from './weather-icon/weather-icon.component';
import { ApiService } from './api/api.service';
import { XmlParsePipe } from './pipes/xmlparse.pipe';
import { TwitterComponent } from './twitter/twitter.component';
import { TwitterService } from './twitter/twitter.service';
import { LogoComponent } from './logo/logo.component';
import { ForecastComponent } from './forecast/forecast.component';
import { TraficComponent } from './trafic/trafic.component';
import { LunchplaceModule } from './lunchplace/lunchplace.module';
import { HumeurComponent } from './humeur/humeur.component';
import { GithubComponent } from './github/github.component';
import { SortPipe } from './pipes/sort.pipe';
import { AnniversaireComponent } from './anniversaire/anniversaire.component';


@NgModule({
  declarations: [
    AppComponent,
    MessageComponent,
    VlilleComponent,
    FirstletterupperPipe,
    ClockComponent,
    HomeComponent,
    SocialComponent,
    CalendarComponent,
    WeatherComponent,
    WeatherIconComponent,
    XmlParsePipe,
    TwitterComponent,
    LogoComponent,
    ForecastComponent,
    TraficComponent,
    HumeurComponent,
    SortPipe,
    AnniversaireComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    GithubModule,
    DateTimeModule,
    LunchplaceModule
  ],
  providers: [
    ApiService,
    VlilleService,
    WeatherService,
    CalendarService,
    TwitterService
  ],
  bootstrap: [AppComponent],
  schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
})
export class AppModule { }
