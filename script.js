/* eslint-disable no-undef */
import { sleep } from 'k6';
import { config } from './config.js';
import { clone, } from './helpers/generator.js';
import exec from 'k6/execution';
import { TestLogin } from './testCases/authLogin.js';
import { TestRegister } from './testCases/authRegister.js';
import { TestNurseManagementPost } from './testCases/nurseManagementPost.js';
import { TestNurseManagementGet } from './testCases/nurseManagementGet.js';
import { TestNurseManagementPut } from './testCases/nurseManagementPut.js';
import { TestNurseManagementDelete } from './testCases/nurseManagementDelete.js';
// import sql from 'k6/x/sql';

const stages = []

if (config.LOAD_TEST) {
    stages.push(
        { target: 50, iterations: 1, duration: "5s" },
        { target: 100, iterations: 1, duration: "10s" },
        { target: 150, iterations: 1, duration: "20s" },
        { target: 200, iterations: 1, duration: "20s" },
        { target: 250, iterations: 1, duration: "20s" },
        { target: 300, iterations: 1, duration: "20s" },
        { target: 600, iterations: 1, duration: "20s" },
        { target: 1200, iterations: 1, duration: "20s" },
    );
} else {
    stages.push({
        target: 1,
        iterations: 1
    });
}

// function determineStage() {
//     let elapsedTime = (exec.instance.currentTestRunDuration / 1000).toFixed(0);
//     if (elapsedTime < 5) return 1; // First 5 seconds
//     if (elapsedTime < 15) return 2; // Next 10 seconds
//     if (elapsedTime < 35) return 3; // Next 20 seconds
//     if (elapsedTime < 55) return 4; // Next 20 secondsd
//     return 5; // Remaining time
// }

export const options = {
    stages,
};

// const positiveCaseConfig = Object.assign(clone(config), {
//     POSITIVE_CASE: true
// })

// function calculatePercentage(percentage, __VU) {
//     return (__VU - 1) % Math.ceil(__VU / Math.round(__VU * percentage)) === 0;
// }


export default function () {
    let tags = {}

    // if (config.LOAD_TEST) { } else {
    const usr = TestRegister(config, tags)
    if (usr) {
        TestLogin(usr, config, tags)
        TestNurseManagementPost(config, usr, tags)
        TestNurseManagementGet(config, usr, tags)
        TestNurseManagementPut(config, usr, tags)
        TestNurseManagementDelete(config, usr, tags)
    }
    sleep(1)
    // }
}
