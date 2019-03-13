<template>
    <v-layout row wrap>
      <v-flex md4>
        <v-list>
          <v-list-item>
            <v-list-tile> 
              <v-list-tile-title class="headline">TODO</v-list-tile-title>
            </v-list-tile>
          </v-list-item>  
        </v-list>
        <v-list>
          <v-list-item>
            <v-card v-for="item in todoItems" v-bind:key="item.id">
              <v-card-title primary-title>
                <div>
                  <h3 class="headline mb-0">{{ item.title }}</h3>
                  <div>{{ item.content }}</div>
                </div>
              </v-card-title>
              <!--  
              <v-card-actions>
                <v-btn flat color="orange">Share</v-btn>
                <v-btn flat color="orange">Explore</v-btn>
              </v-card-actions>
              -->
            </v-card>
          </v-list-item>  
        </v-list>
      </v-flex>

      <v-flex md4>
        <v-list>
          <v-list-item>
            <v-list-tile> 
              <v-list-tile-title class="headline">PROGRESS</v-list-tile-title>
            </v-list-tile>
          </v-list-item>
        </v-list>
        <v-list>
          <v-list-item>
            <v-card v-for="item in progressItems" v-bind:key="item.id">
              <v-card-title primary-title>
                <div>
                  <h3 class="headline mb-0">{{ item.title }}</h3>
                  <div>{{ item.content }}</div>
                </div>
              </v-card-title>
              <!--  
              <v-card-actions>
                <v-btn flat color="orange">Share</v-btn>
                <v-btn flat color="orange">Explore</v-btn>
              </v-card-actions>
              -->
            </v-card>
          </v-list-item>  
        </v-list>
      </v-flex>

      <v-flex md4>
        <v-list>
          <v-list-item>
            <v-list-tile> 
              <v-list-tile-title class="headline">DONE</v-list-tile-title>
            </v-list-tile>
          </v-list-item>
        </v-list>
        <v-list>
          <v-list-item>
            <v-card v-for="item in doneItems" v-bind:key="item.id">
              <v-card-title primary-title>
                <div>
                  <h3 class="headline mb-0">{{ item.title }}</h3>
                  <div>{{ item.content }}</div>
                </div>
              </v-card-title>
              <!--  
              <v-card-actions>
                <v-btn flat color="orange">Share</v-btn>
                <v-btn flat color="orange">Explore</v-btn>
              </v-card-actions>
              -->
            </v-card>
          </v-list-item>  
        </v-list>
      </v-flex>
    </v-layout>

</template>

<script>
import axios from 'axios'
let todoApiUrl = process.env.TODO_API_URL || 'http://localhost:8000'
export default {
  async asyncData (context) {
    const { data: todoItems } = await axios.get(`${todoApiUrl}/todo?status=TODO`)
    const { data: progressItems } = await axios.get(`${todoApiUrl}/todo?status=PROGRESS`)
    const { data: doneItems } = await axios.get(`${todoApiUrl}/todo?status=DONE`)
    return {
      todoItems,
      progressItems,
      doneItems
    }
  }
}
</script>


<style>
.container {
  min-height: 100vh;
  display: flex;
  padding: 2vh;
}
.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}
.state-box {
  float: left;
}
.todotitle {
  font-weight: 300;
  font-size: 20px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}
.links {
  padding-top: 15px;
}
.card {
  border-radius: 5px;
  box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
  transition: 0.3s;
  background-color: #ccccff;
  margin: 10px;
}
.card-title {
  margin-bottom: .75rem;
}
.card:hover {
  box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
}
.card-container {
  padding: 2px 16px;
}
</style>

