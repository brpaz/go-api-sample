__NUXT_JSONP__("/installation-guide", (function(a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,A,B,C,D,E,F,G,H,I,J,K,L,M,N){return {data:[{document:{title:"Installation guide",position:m,category:"Getting started",fullscreen:false,toc:[{id:v,depth:m,text:w},{id:x,depth:m,text:y},{id:z,depth:m,text:A},{id:B,depth:m,text:C},{id:D,depth:E,text:F},{id:G,depth:E,text:H}],body:{type:"root",children:[{type:b,tag:n,props:{id:v},children:[{type:b,tag:h,props:{href:"#pre-requisites",ariaHidden:i,tabIndex:j},children:[{type:b,tag:f,props:{className:[k,l]},children:[]}]},{type:a,value:w}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"Before starting, make sure you have the following tools installed on your machine:"}]},{type:a,value:c},{type:b,tag:I,props:{},children:[{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Git"}]},{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Docker"}]},{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Docker-Compose"}]},{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Make"}]},{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Python (For pre-commit)"}]},{type:a,value:c}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"This project ships a Docker-compose file and a Makefile with common tasks, to facilitate the development process."}]},{type:a,value:c},{type:b,tag:n,props:{id:x},children:[{type:b,tag:h,props:{href:"#setup",ariaHidden:i,tabIndex:j},children:[{type:b,tag:f,props:{className:[k,l]},children:[]}]},{type:a,value:y}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"The first step is clone the project using git to a directory of your choice."}]},{type:a,value:c},{type:b,tag:p,props:{className:[q]},children:[{type:b,tag:r,props:{className:[s,t]},children:[{type:b,tag:e,props:{},children:[{type:b,tag:f,props:{className:[o,u]},children:[{type:a,value:"git"}]},{type:a,value:" clone https:\u002F\u002Fgithub.com\u002Fbrpaz\u002Fgo-api-sample\n"},{type:b,tag:f,props:{className:[o,"builtin","class-name"]},children:[{type:a,value:"cd"}]},{type:a,value:" go-api-sample\n"}]}]}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"After that, run:"}]},{type:a,value:c},{type:b,tag:p,props:{className:[q]},children:[{type:b,tag:r,props:{className:[s,t]},children:[{type:b,tag:e,props:{},children:[{type:b,tag:f,props:{className:[o,u]},children:[{type:a,value:J}]},{type:a,value:" setup\n"}]}]}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"This will run some setup tasks like installing "},{type:b,tag:h,props:{href:"https:\u002F\u002Fpre-commit.com",rel:["nofollow","noopener","noreferrer"],target:"_blank"},children:[{type:a,value:"Pre-commit hooks"}]}]},{type:a,value:c},{type:b,tag:n,props:{id:z},children:[{type:b,tag:h,props:{href:"#launching-the-application",ariaHidden:i,tabIndex:j},children:[{type:b,tag:f,props:{className:[k,l]},children:[]}]},{type:a,value:A}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"To start the application, just run:"}]},{type:a,value:c},{type:b,tag:p,props:{className:[q]},children:[{type:b,tag:r,props:{className:[s,t]},children:[{type:b,tag:e,props:{},children:[{type:b,tag:f,props:{className:[o,u]},children:[{type:a,value:J}]},{type:a,value:" up\n"}]}]}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"This make task will run docker-compose under the hood to build and launch the application containers. It will launch 3 containers:"}]},{type:a,value:c},{type:b,tag:I,props:{},children:[{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"App - With the application code"}]},{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Postgres - for the database"}]},{type:a,value:c},{type:b,tag:g,props:{},children:[{type:a,value:"Nginx-Proxy - for gateway"}]},{type:a,value:c}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"You can run "},{type:b,tag:e,props:{},children:[{type:a,value:"m̀ake logs"}]},{type:a,value:" to see the application logs."}]},{type:a,value:c},{type:b,tag:n,props:{id:B},children:[{type:b,tag:h,props:{href:"#access-the-application",ariaHidden:i,tabIndex:j},children:[{type:b,tag:f,props:{className:[k,l]},children:[]}]},{type:a,value:C}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"There are two ways you can access the application, directly by the container IP address or by domain name."}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"In a more complex project with lots of services, it´s recommended using the domain name approach."}]},{type:a,value:c},{type:b,tag:K,props:{id:D},children:[{type:b,tag:h,props:{href:"#access-by-domain-name",ariaHidden:i,tabIndex:j},children:[{type:b,tag:f,props:{className:[k,l]},children:[]}]},{type:a,value:F}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"By default, the application will be listening on "},{type:b,tag:e,props:{},children:[{type:a,value:"go-api.docker"}]},{type:a,value:" domain.\nYou must add an entry to your "},{type:b,tag:e,props:{},children:[{type:a,value:"\u002Fetc\u002Fhosts"}]},{type:a,value:" pointing this domain to "},{type:b,tag:e,props:{},children:[{type:a,value:L}]},{type:a,value:"  or configure a DNS server to\npoint the domain to "},{type:b,tag:e,props:{},children:[{type:a,value:L}]},{type:a,value:M}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"After that, you can access the application with the following url:"}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:b,tag:e,props:{},children:[{type:a,value:"http:\u002F\u002Fgo-api.docker:8080\u002F_health"}]}]},{type:a,value:c},{type:b,tag:"alert",props:{},children:[{type:a,value:"\n8080 is the port that Nginx Proxy is configured to listen to. The internal application port is not exposed.\n"},{type:b,tag:d,props:{},children:[{type:a,value:"You can change this port, by specifying "},{type:b,tag:"strong",props:{},children:[{type:a,value:"NGINX_PROXY_PORT"}]},{type:a,value:" in the .env file."}]},{type:a,value:c}]},{type:a,value:c},{type:b,tag:K,props:{id:G},children:[{type:b,tag:h,props:{href:"#access-by-container-ip",ariaHidden:i,tabIndex:j},children:[{type:b,tag:f,props:{className:[k,l]},children:[]}]},{type:a,value:H}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"You can find the container IP, by running "},{type:b,tag:e,props:{},children:[{type:a,value:"docker-compose ps"}]},{type:a,value:" and inspecting the \"Ports\" section of the output."}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"For example, "},{type:b,tag:e,props:{},children:[{type:a,value:"0.0.0.0:32768-\u003E5000\u002Ftcp"}]},{type:a,value:" indicates the application is listening on port 32768 as the internal port is never exposed."}]},{type:a,value:c},{type:b,tag:d,props:{},children:[{type:a,value:"You can then access the application, using "},{type:b,tag:e,props:{},children:[{type:a,value:"http:\u002F\u002Flocalhost:32768"}]},{type:a,value:M}]}]},dir:"\u002Fen",path:"\u002Fen\u002Finstallation-guide",extension:".md",slug:"installation-guide",createdAt:N,updatedAt:N,to:"\u002Finstallation-guide"},prev:{title:"Directory Structure",slug:"directory-structure",to:"\u002Fdirectory-structure"},next:{title:"The Makefile",slug:"the-makefile",to:"\u002Fthe-makefile"}}],fetch:[],mutations:[]}}("text","element","\n","p","code","span","li","a","true",-1,"icon","icon-link",2,"h2","token","div","nuxt-content-highlight","pre","language-shell","line-numbers","function","pre-requisites","Pre-Requisites","setup","Setup","launching-the-application","Launching the application","access-the-application","Access the application.","access-by-domain-name",3,"Access by domain name.","access-by-container-ip","Access by container IP","ul","make","h3","127.0.0.1",".","2020-10-17T15:15:06.787Z")));