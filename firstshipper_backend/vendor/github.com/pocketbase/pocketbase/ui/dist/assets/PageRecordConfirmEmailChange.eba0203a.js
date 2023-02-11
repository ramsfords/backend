import{S as z,i as A,s as G,F as I,c as T,m as L,t as v,a as y,d as R,C as J,E as M,g as _,k as N,n as W,o as b,G as Y,H as j,p as B,q as D,e as m,w as C,b as h,f as d,r as P,h as k,u as q,v as K,y as E,x as O,z as F}from"./index.536fddcb.js";function Q(r){let e,t,s,l,n,o,c,i,a,u,g,$,p=r[3]&&S(r);return o=new D({props:{class:"form-field required",name:"password",$$slots:{default:[V,({uniqueId:f})=>({8:f}),({uniqueId:f})=>f?256:0]},$$scope:{ctx:r}}}),{c(){e=m("form"),t=m("div"),s=m("h5"),l=C(`Type your password to confirm changing your email address
                    `),p&&p.c(),n=h(),T(o.$$.fragment),c=h(),i=m("button"),a=m("span"),a.textContent="Confirm new email",d(t,"class","content txt-center m-b-base"),d(a,"class","txt"),d(i,"type","submit"),d(i,"class","btn btn-lg btn-block"),i.disabled=r[1],P(i,"btn-loading",r[1])},m(f,w){_(f,e,w),k(e,t),k(t,s),k(s,l),p&&p.m(s,null),k(e,n),L(o,e,null),k(e,c),k(e,i),k(i,a),u=!0,g||($=q(e,"submit",K(r[4])),g=!0)},p(f,w){f[3]?p?p.p(f,w):(p=S(f),p.c(),p.m(s,null)):p&&(p.d(1),p=null);const H={};w&769&&(H.$$scope={dirty:w,ctx:f}),o.$set(H),(!u||w&2)&&(i.disabled=f[1]),(!u||w&2)&&P(i,"btn-loading",f[1])},i(f){u||(v(o.$$.fragment,f),u=!0)},o(f){y(o.$$.fragment,f),u=!1},d(f){f&&b(e),p&&p.d(),R(o),g=!1,$()}}}function U(r){let e,t,s,l,n;return{c(){e=m("div"),e.innerHTML=`<div class="icon"><i class="ri-checkbox-circle-line"></i></div> 
            <div class="content txt-bold"><p>Successfully changed the user email address.</p> 
                <p>You can now sign in with your new email address.</p></div>`,t=h(),s=m("button"),s.textContent="Close",d(e,"class","alert alert-success"),d(s,"type","button"),d(s,"class","btn btn-secondary btn-block")},m(o,c){_(o,e,c),_(o,t,c),_(o,s,c),l||(n=q(s,"click",r[6]),l=!0)},p:E,i:E,o:E,d(o){o&&b(e),o&&b(t),o&&b(s),l=!1,n()}}}function S(r){let e,t,s;return{c(){e=C("to "),t=m("strong"),s=C(r[3]),d(t,"class","txt-nowrap")},m(l,n){_(l,e,n),_(l,t,n),k(t,s)},p(l,n){n&8&&O(s,l[3])},d(l){l&&b(e),l&&b(t)}}}function V(r){let e,t,s,l,n,o,c,i;return{c(){e=m("label"),t=C("Password"),l=h(),n=m("input"),d(e,"for",s=r[8]),d(n,"type","password"),d(n,"id",o=r[8]),n.required=!0,n.autofocus=!0},m(a,u){_(a,e,u),k(e,t),_(a,l,u),_(a,n,u),F(n,r[0]),n.focus(),c||(i=q(n,"input",r[7]),c=!0)},p(a,u){u&256&&s!==(s=a[8])&&d(e,"for",s),u&256&&o!==(o=a[8])&&d(n,"id",o),u&1&&n.value!==a[0]&&F(n,a[0])},d(a){a&&b(e),a&&b(l),a&&b(n),c=!1,i()}}}function X(r){let e,t,s,l;const n=[U,Q],o=[];function c(i,a){return i[2]?0:1}return e=c(r),t=o[e]=n[e](r),{c(){t.c(),s=M()},m(i,a){o[e].m(i,a),_(i,s,a),l=!0},p(i,a){let u=e;e=c(i),e===u?o[e].p(i,a):(N(),y(o[u],1,1,()=>{o[u]=null}),W(),t=o[e],t?t.p(i,a):(t=o[e]=n[e](i),t.c()),v(t,1),t.m(s.parentNode,s))},i(i){l||(v(t),l=!0)},o(i){y(t),l=!1},d(i){o[e].d(i),i&&b(s)}}}function Z(r){let e,t;return e=new I({props:{nobranding:!0,$$slots:{default:[X]},$$scope:{ctx:r}}}),{c(){T(e.$$.fragment)},m(s,l){L(e,s,l),t=!0},p(s,[l]){const n={};l&527&&(n.$$scope={dirty:l,ctx:s}),e.$set(n)},i(s){t||(v(e.$$.fragment,s),t=!0)},o(s){y(e.$$.fragment,s),t=!1},d(s){R(e,s)}}}function x(r,e,t){let s,{params:l}=e,n="",o=!1,c=!1;async function i(){if(o)return;t(1,o=!0);const g=new Y("../");try{const $=j(l==null?void 0:l.token);await g.collection($.collectionId).confirmEmailChange(l==null?void 0:l.token,n),t(2,c=!0)}catch($){B.errorResponseHandler($)}t(1,o=!1)}const a=()=>window.close();function u(){n=this.value,t(0,n)}return r.$$set=g=>{"params"in g&&t(5,l=g.params)},r.$$.update=()=>{r.$$.dirty&32&&t(3,s=J.getJWTPayload(l==null?void 0:l.token).newEmail||"")},[n,o,c,s,i,l,a,u]}class te extends z{constructor(e){super(),A(this,e,x,Z,G,{params:5})}}export{te as default};
