<template>
  <div id="app">
  	<h1 class="logo" style="margin-bottom:"><span>Software</span> <span>Lab</span></h1>
      <software-list v-on:changeBundle="changeBundle" :openTab="openTab"></software-list>
  </div>
</template>

<style lang="scss">
  @import url(https://db.onlinewebfonts.com/c/9126c154a82d08fe6124045ce0fc61d9?family=AlmaqW01-Refined);

  $gray: #1e262c;
  $yellow: #f6c435;
  $orange: #e16622;
  $red: #d11d29;
  $purple: #542437;
  $blue: #4B75BD;

  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  html {
    background-color: #141a1e;
  }

  p {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
  }

  a {
    text-decoration: none;
    color: #2c3e50;
  }

  h1 {
  	font-size: 4rem;
  }

  h1, h2 {
    font-family: 'https://db.onlinewebfonts.com/c/9126c154a82d08fe6124045ce0fc61d9?family=AlmaqW01-Refined', 'Avenir', Helvetica, Arial, sans-serif;
    font-weight: 500;
  }

  h3, h4 {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    font-weight: 500;
  }

  h3 {
  	font-size: 2rem;
  }

  h4 {
  	font-size: 1.5rem;
  }

  ul {
    list-style-type: none;
  }

  .logo {
  	position: absolute;
  	top: 20px;
  	right: 20px;
  	width: 100%;
  	text-align: right;
  	padding-right: 20px;
  	color: #222;
  	font-size: 2rem;

  	span:first-of-type {
  		color: $yellow;
  	}
  	span:nth-of-type(2) {
  		color: $yellow;
  	}
  	span:nth-of-type(3) {
  		color: $orange;
  	}
  	span:nth-of-type(4) {
  		color: $blue;
  	}
  }

  .big-input {
  	position: relative;
  	background-color: transparent;
  	border: none;
  	outline: none;

  	font-family: 'https://db.onlinewebfonts.com/c/9126c154a82d08fe6124045ce0fc61d9?family=AlmaqW01-Refined', 'Avenir', Helvetica, Arial, sans-serif;
  	font-size: 2.5rem;
  	color: #ddd5bf;
  	text-align: right;
  }

  .big-input-line {
  	content: " ";
  	width: 0;
  	height: 2px;
  	position: absolute;
  	right: 0;
  	bottom: 0;
  	background-color:red;

  	overflow: visible;
  }

  .big-input:focus {
  	& ~ big-input-line {
  		width: 100%;
  	}
  }
</style>

<script>
import SoftwareList from './components/SoftwareList.vue'
import BundleList from './components/BundleList.vue'

export default {
  name: 'app',
  components: {
    SoftwareList,
    BundleList
  },

  data () {
  	return {
  		openTab: ''
  	}
  },

  methods: {
    changeBundle: function (name) {
      this.openTab = name;
    }
  }
}
</script>
