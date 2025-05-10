<script>
export default {
    data: function(){
        return{
            errormsg: null,
            departments: [],
            selectedDepartment: "",
            newAddress: "",
            isShow: false,
        };
    },
    mounted() {
        this.getDepartments();
    },
    methods: {
        async getDepartments() {
            this.errormsg = "";
            try {
                let response = await this.$axios.get(
                    `profiles/${sessionStorage.username}/departments`,
                    { headers: { Session: sessionStorage.session, Token: sessionStorage.token } }
                );
                this.departments = response.data || [];
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async modifyAddress(){
            this.errormsg = "";
            try {
                let response = await this.$axios.put(
                    `profiles/${sessionStorage.username}/departments/${this.selectedDepartment}`,
                    { address: this.newAddress },
                    { headers: { Session: sessionStorage.session, Token: sessionStorage.token } }
                );
                this.getDepartments();
                this.isShow = false;
                this.newAddress = "";
                this.selectedDepartment = "";
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        changeShow(department){
            this.isShow = !this.isShow;
            this.selectedDepartment = department;
        },
    },
}
</script>

<template>
    <div class="department-conteiner">
        <h1 class="title">Departments</h1>
        <hr class="divider" />

        <div v-if="errormsg" class="error">{{ errormsg }}</div>

        <div class="department-list">
            <div
                v-for="department in departments"
                class="department-box"
            >
                <div class="department-info">
                    <h2 class="department-title">{{ department.name }}</h2>
                    <p><strong>ID:</strong> {{ department.id }}</p>
                    <p><strong>Manager:</strong> {{ department.manager }}</p>
                    <p><strong>Address:</strong> {{ department.address }}</p>
                </div>

                <div class="button-container">
                    <div v-if="isShow">
                        <input
                            v-model="newAddress"
                            type="text"
                            placeholder="New Address"
                            class="address-input"
                        />
                        <button class="modify-button" @click="modifyAddress">Modify Address</button>
                    </div>
                    <div v-else>
                        <button class="modify-button" @click="changeShow(department.name)">Modify Address</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.department-conteiner {
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

.department-list {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.department-box {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    border: 1px solid #ccc;
    border-radius: 12px;
    padding: 20px;
    background-color: #f9f9f9;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.05);
}

.department-info {
    flex-grow: 1;
}

.department-title {
    font-size: 20px;
    font-weight: bold;
    margin-bottom: 10px;
}

.button-container {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    padding-left: 20px;
}

.address-input {
    width: 100%;
    padding: 10px 15px;
    border: 1px solid #ccc;
    border-radius: 8px;
    font-size: 16px;
    outline: none;
    box-shadow: none;
}

.modify-button {
    padding: 8px 14px;
    background-color: #12B886;
    border: none;
    border-radius: 8px;
    color: white;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
}

.error {
    color: red;
    text-align: center;
    margin-top: 10px;
}

</style>