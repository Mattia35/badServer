<script>
export default {
    data: function(){
        return{
            errormsg: null,
            projects: [],
            selectedProject: "",
        };
    },
    methods: {
        async getProjects() {
            this.errormsg = "";
            if (this.selectedProject.length === 0) {
                this.projects = [];
                return;
            }
            try {
                console.log(sessionStorage.session, sessionStorage.token);
                let response = await this.$axios.get(
                    `profiles/${sessionStorage.username}/projects?name=${this.selectedProject}`,
                    { headers: { Session: sessionStorage.session, Token: sessionStorage.token } }
                );
                this.projects = response.data || [];
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
    }
}
</script>

<template>
    <div class="projects-container">
        <h1 class="title">Projects</h1>
        <hr class="divider" />
  
        <div class="search-bar">
            <input
            v-model="selectedProject"
            type="text"
            placeholder="Search projects..."
            class="search-input"
            />
            <button class="search-button" @click="getProjects">Search</button>
        </div>
    
        <div v-if="errormsg" class="error">{{ errormsg }}</div>

        <div class="projects-list">
            <div
                v-for="project in projects"
                class="project-box"
            >
                <h2 class="project-title">{{ project.project.name }}</h2>
                <p><strong>ID:</strong> {{ project.project.id }}</p>
                <p><strong>Start Date:</strong> {{ project.project.start_date }}</p>
                <p><strong>End Date:</strong> {{ project.project.end_date }}</p>
                <p><strong>Department:</strong> {{ project.project.department }}</p>
                <p><strong>Employees:</strong></p>
                <div v-for="employee in project.employees">
                    <p>{{ employee }}</p>
                </div>
            </div>
        </div>
    </div>
</template>  

<style>
.projects-container {
    width: 100%;
    margin: 0 auto;
    padding: 0 20px;
    display: flex;
    flex-direction: column;        
    justify-content: flex-start; 
    height: 100%;
}

.title {
    font-size: 32px;
    font-weight: bold;
    text-align: center;
    margin-bottom: 10px;
    padding-top: 0;
}

.divider {
    border: none;
    border-top: 2px solid #ccc;
    margin-bottom: 20px;
    width: 100%;
}

.search-bar {
    width: 100%;
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-bottom: 20px;
    max-width: 800px;
    margin-left: auto;
    margin-right: auto;
}

.search-input {
    width: 100%;
    padding: 10px 15px;
    border: 1px solid #ccc;
    border-radius: 8px;
    font-size: 16px;
    outline: none;
    box-shadow: none;
}

.search-input:focus {
    outline: none;
    box-shadow: none;
    border-color: #ccc; 
}

.search-button {
    padding: 10px 15px;
    background-color: #12B886;
    border: none;
    border-radius: 8px;
    color: white;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
}

.projects-list {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.project-box {
    border: 1px solid #ccc;
    border-radius: 12px;
    padding: 20px;
    background-color: #f9f9f9;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.05);
}

.project-title {
    font-size: 20px;
    font-weight: bold;
    margin-bottom: 10px;
}

.error {
    color: red;
    text-align: center;
    margin-top: 10px;
}

</style>
