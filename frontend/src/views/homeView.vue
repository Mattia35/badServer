<script>
export default{
    data: function(){
        return{
            errormsg: null,
        };
    },
    methods: {
        async goToEmployees(){
            this.errormsg = null;
            try {
                // redirect to the employees page
                this.$router.push(`/${sessionStorage.username}/employees`);
            } catch (e) {
                this.errormsg = e.toString();
            };
        },
        async goToProjects(){
            this.errormsg = null;
            try {
                // redirect to the projects page
                this.$router.push(`/${sessionStorage.username}/projects`);
            } catch (e) {
                this.errormsg = e.toString();
            };
        },
        async goToDepartments(){
            this.errormsg = null;
            try {
                // redirect to the departments page
                this.$router.push(`/${sessionStorage.username}/departments`);
            } catch (e) {
                this.errormsg = e.toString();
            };
        },
    }
}
</script>

<template>
    <div class="container">
      <div v-if="errormsg" class="error">{{ errormsg }}</div>
  
      <div class="card" @click="goToEmployees">
        <div class="card-content">
          <h2>Employees</h2>
        </div>
      </div>
  
      <div class="card" @click="goToProjects">
        <div class="card-content">
          <h2>Projects</h2>
        </div>
      </div>
  
      <div class="card" @click="goToDepartments">
        <div class="card-content">
          <h2>Departments</h2>
        </div>
      </div>
    </div>
  </template>

<style>
.container {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    padding: 20px;
    width: 1200px;
    flex-wrap: wrap;
}

.card {
    display: flex;
    align-items: center;
    background-color: #f5f5f5;
    border-radius: 12px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 100%;
    padding: 15px;
    margin: 10px;
    cursor: pointer;
    transition: transform 0.2s;
}

.card:hover {
    transform: scale(1.03);
}

.card-image {
    width: 60px;
    height: 60px;
    margin-right: 15px;
}

.card-content {
    flex-grow: 1;
}

.card-button {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 8px 12px;
    border-radius: 6px;
    cursor: pointer;
    margin-top: 10px;
}

.card-button:hover {
    background-color: #0056b3;
}

.error {
    color: red;
    margin-bottom: 10px;
    width: 100%;
    text-align: center;
}
</style>