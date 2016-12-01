import { GithubService } from './github.service';
import { GithubComponent } from './github.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@NgModule({
    declarations: [ GithubComponent ],
    exports: [ GithubComponent ],
    providers: [ GithubService],
    imports: [ CommonModule ]
})
export class GithubModule {

}