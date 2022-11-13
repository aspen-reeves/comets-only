import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Profile } from './profile.model';
import { Profile2 } from './profile2.model';
import * as EventEmitter from 'events';
import { Router } from '@angular/router';
import { Subject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class APIService {
  subject$ = new Subject<Profile[]>();

  profile: Profile2 = new Profile2();
  aspenAPI: string = `http://144.126.154.126:10000/`;

  profiles: Profile[] = [];
  userId: any;
  signedIn: boolean;

  constructor(private httpClient: HttpClient, public router: Router) {}

  // SIGN IN REQUEST
  signin(credentials: { u: string; p: string }) {
    this.httpClient
      .post(`${this.aspenAPI}auth`, JSON.stringify(credentials))
      .subscribe((res) => {
        console.log(res);

        if (res != 'wrong password' && res > 0) { // SUCCESS
          this.userId = res;
          this.signedIn = true;
          this.router.navigate(['menu']);
        } else if (res == 0) {
          this.router.navigate(['signup']);
        } else this.signedIn = false;
      });
  }

  // GET PROFILE REQUEST
  getProfile() {
    this.httpClient
      .get<Profile>('http://144.126.154.126:10000/getbitches')
      .subscribe((data: Profile) => {
        let profile = { ...data };
        this.profiles.push(profile);

        // Change Array through multicast
        this.subject$.next(this.profiles);
      });
  }

  // **IF JSON KEYS WERE CAPITALIZED**
  // newProfileWindow(someProfile: Profile): Profile {
  //   console.log(someProfile);

  //   return {
  //     name: someProfile.name,
  //     age: someProfile.age,
  //     lang: someProfile.lang,
  //     os: someProfile.os,
  //     editor: someProfile.editor,
  //     lastShower: someProfile.lastShower,
  //     code: someProfile.code,
  //   };
  // }
}
