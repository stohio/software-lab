<template>

  <div id="list-container">
    <ul 
    class="server-list">
      <li class="server">
        <div class="flex-container" v-bind:class="{hidden: expandList === false}" style="margin:auto">
          <h3 style="margin:auto">Available Software to Build With</h3>
          <input v-model="searchString" type="text" placeholder="Search..." class="big-input" style="margin-bottom:20px"><span class="big-input-line"></span>
        </div>
        <ul class="software-list" v-bind:class="{hidden: expandList === false}">
          <div v-for="software in searchFilter(server.softwares)" v-on:click="downloadFile(software)">
            <li class="software-card">
              <h4>{{ software.name }} - {{ software.os }}</h4>
              <div class="info">
              <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
              width="42px" height="32px" viewBox="0 0 42 32" enable-background="new 0 0 42 32" xml:space="preserve">
              <defs>
                <linearGradient id="soft-gradient" gradientTransform="rotate(60)">
                  <stop offset="0%" stop-color="#d11d29"></stop>
                  <stop offset="100%" stop-color="#542437"></stop>
                </linearGradient>
              </defs>
              <g>
                <path d="M33.958,12.982C33.528,6.372,28.931,0,20.5,0c-1.029,0-2.044,0.1-3.018,0.297
                c-0.271,0.055-0.445,0.318-0.391,0.59c0.055,0.271,0.314,0.445,0.59,0.391C18.589,1.093,19.538,1,20.5,1C29.088,1,33,7.739,33,14
                v1.5c0,0.276,0.224,0.5,0.5,0.5s0.5-0.224,0.5-0.5V14c0-0.005-0.001-0.011-0.001-0.016C37.062,14.248,41,16.916,41,22.5
                c0,4.767-3.514,8.5-8,8.5H9c-3.976,0-8-2.92-8-8.5C1,18.406,3.504,14,9,14h1.5c0.276,0,0.5-0.224,0.5-0.5S10.776,13,10.5,13H9v-2
                c0-3.727,2.299-6.042,6-6.042c3.364,0,6,2.654,6,6.042v12.993l-4.16-3.86c-0.2-0.188-0.517-0.177-0.706,0.026
                c-0.188,0.202-0.177,0.519,0.026,0.706l4.516,4.189c0.299,0.298,0.563,0.445,0.827,0.445c0.261,0,0.52-0.145,0.808-0.433
                l4.529-4.202c0.203-0.188,0.215-0.504,0.026-0.706c-0.188-0.204-0.506-0.215-0.706-0.026L22,23.993V11c0-3.949-3.075-7.042-7-7.042
                c-4.252,0-7,2.764-7,7.042v2.051c-5.255,0.508-8,5.003-8,9.449C0,27.105,3.154,32,9,32h24c5.047,0,9-4.173,9-9.5
                C42,16.196,37.443,13.222,33.958,12.982z"/>
              </g>
            </svg>
              <p>Download</p>
              </div>
            </li>
          </div>
        </ul>
      </li>
    </ul>
    <div class="bundle-card see-all"
      v-on:click="expandListItems"
      v-bind:class="{hidden: expandList}">
        <div class="block"></div>
        <svg viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" fill="#000"><g><path d="M 23.706,15.312L 17.788,8.622c-0.392-0.392-1.030-0.392-1.422,0s-0.392,1.030,0,1.422l 5.3,5.992l-5.3,5.992 c-0.392,0.392-0.392,1.030,0,1.422s 1.030,0.392, 1.422,0l 5.918-6.69c 0.2-0.2, 0.296-0.462, 0.292-0.724 C 24,15.774, 23.906,15.512, 23.706,15.312zM 9.192,23.452c 0.392,0.392, 1.030,0.392, 1.422,0l 5.918-6.69c 0.2-0.2, 0.296-0.462, 0.292-0.724 c 0.004-0.262-0.092-0.526-0.292-0.724L 10.616,8.622c-0.392-0.392-1.030-0.392-1.422,0s-0.392,1.030,0,1.422l 5.3,5.992l-5.3,5.992 C 8.8,22.422, 8.8,23.060, 9.192,23.452z"></path></g></svg>
        <h2>See All Software</h2>
    </div>

  </div>
</template>

<script>
export default {
  name: 'software-list',
  data () {
    return {     
apiUrl:  "http://" + window.location.hostname + ":8080", //NOT USED
      searchString: '',
      expandList: false,
        
      server: {
      },    
    }
  },

  props: ['openTab'],

  beforeMount () {
      this.fetchServerList();
      this.expandListItems();

  },

  watch: {
    openTab: function (val, oldVal) {
      this.expandList = true;
    }
  },

  methods: {
    fetchServerList () {
      var xhr = new XMLHttpRequest();
      var self = this;
      var trustMe = self.apiUrl;
      //trustMe = trustMe.substr(0, trustMe.indexOf('/softwarelab'));
        console.log("About to open networks ", trustMe + "/networks");
      xhr.open('GET', trustMe+"/networks");
      
      xhr.onload = function() {
          self.server.private_ip = xhr.responseText;
          console.log(xhr.responseText);
          
          var url = window.location.href;
          console.log("FUCK");
          
          var response = JSON.parse(xhr.responseText);
          self.server = response[0];
          self.server.softwares = [];
          
          for(var i = 0; i<self.server.stack.softwares.length; i++){
              console.log(self.server.stack.softwares[i]);
              for(var j = 0; j<self.server.stack.softwares[i].versions.length; j++){
                  self.server.softwares.push({
                    name: self.server.stack.softwares[i].name,
                    os:   self.server.stack.softwares[i].versions[j].os,
                      softid: self.server.stack.softwares[i].id,
                      verid: self.server.stack.softwares[i].versions[j].id,
                      architecture: self.server.stack.softwares[i].versions[j].architecture
                      
                  });
              }
          }
        }
        
        xhr.send();
    },
      
    searchFilter(software) {
      var self = this;
      return software.filter(function (software) {

        if (self.searchString === '')
          return software;
        
        
        var searchParams = self.searchString.split(' ');

        for (var i = 0; i < searchParams.length; i++) 
          if (software.name.toLowerCase().includes(searchParams[i].toLowerCase())) 
            return software;
          
        return 0;
      })
    },

      
//Bundles not currently being      
//    bundleFilter(software) {
//      var self = this;
//
//      return software.filter(function (software) {
//
//        if (self.openTab != '') {
//            
//          for (var i = 0; i < self.bundles[self.openTab].length; i ++) {
//            if ( self.bundles[self.openTab][i] === software.id ) {
//              return software;
//            }
//          }
//        } else {
//          return software;
//        }
//
//      })
//    }, 

    downloadFile(software) {
      var self = this;
      var softid = software.softid;
      var verid = software.verid;
      var xhr = new XMLHttpRequest();
        
      var trustMe = self.apiUrl;
      //trustMe = trustMe.substr(0, trustMe.indexOf('/softwarelab'));
        console.log(trustMe);
        
      xhr.open('GET', trustMe+"/software/" + softid + "/versions/" + verid);
      xhr.onload = function() {
          //parse response into JSON and get response.ip
          var response = JSON.parse(xhr.responseText);
          var ip = response.ip;
          console.log("ip: " + ip);
          if(ip != "undefined")
            window.location.href = "http://" + ip + "/download/software/" + softid + "/versions/" + verid;
    
          
      //xhr.open('GET', "http://" + ip + "/download/software/" + softid + "/versions/" + verid);
      //xhr.onload = function() {
          //this should just download it right cam i don't need to put anything here. 
      //}
      //xhr.send();
        }
      xhr.send();
          //request to new url which download software
    },

    expandListItems: function() {
        this.expandList = true;
        this.$emit('changeBundle', '');
    },

    dontShowAll: function() {
      this.expandList = false;
    }
  }
}
</script>

<style scoped lang="scss">
#list-container {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ddd5bf;

  width: 100%;
  max-width: 940px;

  margin: 20px auto 40px;
}

.hidden {
  max-height: 0px;
  overflow: hidden;
  opacity: 0;
}

.server {
  width: 100%;
}

.server-info {
  position: fixed;
  bottom: 0;
  right: -20px;
  opacity: 0.5;
  transform: scale(0.6);
  text-align: right;
}

h3 {
  margin-bottom: 20px;
  text-align: left;
  display: inline-block;
}

.flex-container {
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}

.software-card {
  display: block;
  width: 100%;
  background-color: rgba(10, 13, 16, 0.8);
  color: #ddd5bf;

  text-align: left;

  padding: 20px 40px;
  margin-bottom: 20px;

  cursor: pointer;

  border-radius: 5px;
  box-shadow: 0 1px 4px 0 rgba(0,0,0,0.16);

  transition: 0.3s box-shadow ease;

  &:hover {
    box-shadow: 0 3px 6px 0 rgba(0,0,0,0.23);
  }

  &:hover, &:active {
    
    h2 {
      transform: translateY(-3px);
    }

    .block {
      top: 0%;
      left: 50%;
    }
  }
}

.software-card {

  position: relative;

  .info {
    position: absolute;
    right: 40px;
    top: 10px;
    padding: 10px 0;

    display: flex;
    justify-content: row-end;
    align-items: center;
  }

  p {
    color: #000;

    margin-left: 20px;
    transition: 0.2s color ease;
  }

  svg {
    fill: #000;
    transition: 0.2s fill ease;
  }

  lineargradient {
    stop {
      stop-color: #000;
      transition: 0.2s stop-color ease;
    }

    stop:last-of-type {
      transition-delay: 0s;
    }
  }
}

.software-card:hover {
  p {
    color: #f6c435; 
  }


  svg {
    fill: #f6c435;
  }
  lineargradient {
    stop:first-of-type {
      stop-color: #3a5255;
    }
    stop:last-of-type {
      stop-color: #f6c435;
    }
  }
}
</style>
