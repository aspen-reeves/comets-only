import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Profile } from './profile.model';
import { Profile2 } from './profile2.model';
import * as EventEmitter from 'events';
import { Router } from '@angular/router';
import { Subject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class APIService {
  // USER INFO
  name: string;
  userId: any;

  subject$ = new Subject<Profile>();
  aspenAPI: string = `http://144.126.154.126:10000/`;

  signedIn: boolean;

  constructor(private httpClient: HttpClient, public router: Router) {}

  // SIGN IN REQUEST
  signin(credentials: { u: string; p: string }) {
    this.httpClient
      .post(`${this.aspenAPI}auth`, JSON.stringify(credentials))
      .subscribe(
        (res) => {
          console.debug(res);

          if (res != 'wrong password' && res > 0) {
            // SUCCESS
            this.userId = res;
            this.signedIn = true;
            this.router.navigate(['menu']);
          } else this.signedIn = false;
        },
        (err) => {
          console.log(err);
          if (err.status === 200) {
            this.router.navigate(['signup']);
          }
        }
      );
  }

  // GET PROFILE REQUEST
  getProfile() {
    this.httpClient
      .get<Profile>('http://144.126.154.126:10000/getbitches')
      .subscribe((data: Profile) => {
        let profile = { ...data };

        // Change Array through multicast
        this.subject$.next(profile);
      });
  }

  storeToken() {
    sessionStorage.setItem('token', JSON.stringify(this.userId));
  }

  getToken(): number {
    const token = sessionStorage.getItem('token') as string;
    return JSON.parse(token);
  }
}
