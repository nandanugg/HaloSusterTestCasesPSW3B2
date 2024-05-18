import { fail } from "k6";
import { isEqualWith, isExists, isOrdered, isTotalDataInRange, isValidDate } from "../helpers/assertion.js";
import { testGetAssert } from "../helpers/request.js";
import { isItUserValid, isNurseUserValid } from "../types/user.js";
import { generateRandomNumber } from "../helpers/generator.js";
/**
 * 
 * @param {import("../config.js").Config} config 
 * @param {Object} tags 
 * @param {import("../types/user.js").ItUser} userIt
 * @param {import("../types/user.js").NurseUser}userNurse 
 */
export function TestMedicalRecordGet(config, userIt, userNurse, tags) {
    const currentRoute = `${config.BASE_URL}/v1/medical/record`
    const currentFeature = "medical record get"
    if (!isItUserValid(userIt)) {
        fail(`${currentFeature} Invalid userIt object`)
    }
    if (!isNurseUserValid(userNurse)) {
        fail(`${currentFeature} Invalid userNurse object`)
    }
    /** @type {import("../helpers/request.js").RequestAssertResponse} */
    const headers = {
        Authorization: `Bearer ${generateRandomNumber(0, 1) ? userNurse.accessToken : userIt.accessToken}`
    }

    if (!config.POSITIVE_CASE) {
        testGetAssert(currentFeature, "no header", currentRoute, {}, {}, {
            ['should return 401']: (res) => res.status === 401,
        }, config, tags);
    }

    testGetAssert(currentFeature, "get all record", currentRoute, {}, headers, {
        ['should return 200']: (res) => res.status === 200,
        ['should all have a identityDetail']: (res) => isExists(res, "data[].identityDetail"),
        ['should all have a identityDetail.identityNumber']: (res) => isExists(res, "data[].identityDetail.identityNumber"),
        ['should all have a identityDetail.phoneNumber']: (res) => isExists(res, "data[].identityDetail.phoneNumber"),
        ['should all have a identityDetail.name']: (res) => isExists(res, "data[].identityDetail.name"),
        ['should all have a identityDetail.birthDate and the format should be date']: (res) => isEqualWith(res, 'data[].identityDetail.birthDate', (v) => v.every(a => isValidDate(a))),
        ['should all have a identityDetail.gender']: (res) => isExists(res, "data[].identityDetail.gender"),
        ['should all have a identityDetail.identityCardScanImg']: (res) => isExists(res, "data[].identityDetail.identityCardScanImg"),
        ['should all have a symptoms']: (res) => isExists(res, "data[].symptoms"),
        ['should all have a medications']: (res) => isExists(res, "data[].medications"),
        ['should all have a createdBy']: (res) => isExists(res, "data[].createdBy"),
        ['should all have a createdBy.nip']: (res) => isExists(res, "data[].createdBy.nip"),
        ['should all have a createdBy.name']: (res) => isExists(res, "data[].createdBy.name"),
        ['should all have a createdBy.userId']: (res) => isExists(res, "data[].createdBy.userId"),
        ['should have createdAt and format should be date']: (res) => isEqualWith(res, 'data[].createdAt', (v) => v.every(a => isValidDate(a))),
        ['should not have more than 5 result']: (res) => isTotalDataInRange(res, 'data[]', 1, 5),
    }, config, tags);

    testGetAssert(currentFeature, "get all record with createdBy.userId", currentRoute, { 'createdBy.userid': userNurse.userId }, headers, {
        ['should return 200']: (res) => res.status === 200,
        ['should all have a identityDetail']: (res) => isExists(res, "data[].identityDetail"),
        ['should all have a identityDetail.identityNumber']: (res) => isExists(res, "data[].identityDetail.identityNumber"),
        ['should all have a identityDetail.phoneNumber']: (res) => isExists(res, "data[].identityDetail.phoneNumber"),
        ['should all have a identityDetail.name']: (res) => isExists(res, "data[].identityDetail.name"),
        ['should all have a identityDetail.birthDate and the format should be date']: (res) => isEqualWith(res, 'data[].identityDetail.birthDate', (v) => v.every(a => isValidDate(a))),
        ['should all have a identityDetail.gender']: (res) => isExists(res, "data[].identityDetail.gender"),
        ['should all have a identityDetail.identityCardScanImg']: (res) => isExists(res, "data[].identityDetail.identityCardScanImg"),
        ['should all have a symptoms']: (res) => isExists(res, "data[].symptoms"),
        ['should all have a medications']: (res) => isExists(res, "data[].medications"),
        ['should all have a createdBy']: (res) => isExists(res, "data[].createdBy"),
        ['should all have a createdBy.name']: (res) => isExists(res, "data[].createdBy.name"),
        ['should all have a createdBy.nip']: (res) => isEqualWith(res, "data[].createdBy.nip", userNurse.nip),
        ['should all have a createdBy.userId']: (res) => isEqualWith(res, "data[].createdBy.userId", userNurse.userId),
        ['should have createdAt and format should be date']: (res) => isEqualWith(res, 'data[].createdAt', (v) => v.every(a => isValidDate(a))),
        ['should not have more than 5 result']: (res) => isTotalDataInRange(res, 'data[]', 1, 5),
    }, config, tags);

    testGetAssert(currentFeature, "get all record with nip", currentRoute, { 'nip': userNurse.nip }, headers, {
        ['should return 200']: (res) => res.status === 200,
        ['should all have a identityDetail']: (res) => isExists(res, "data[].identityDetail"),
        ['should all have a identityDetail.identityNumber']: (res) => isExists(res, "data[].identityDetail.identityNumber"),
        ['should all have a identityDetail.phoneNumber']: (res) => isExists(res, "data[].identityDetail.phoneNumber"),
        ['should all have a identityDetail.name']: (res) => isExists(res, "data[].identityDetail.name"),
        ['should all have a identityDetail.birthDate and the format should be date']: (res) => isEqualWith(res, 'data[].identityDetail.birthDate', (v) => v.every(a => isValidDate(a))),
        ['should all have a identityDetail.gender']: (res) => isExists(res, "data[].identityDetail.gender"),
        ['should all have a identityDetail.identityCardScanImg']: (res) => isExists(res, "data[].identityDetail.identityCardScanImg"),
        ['should all have a symptoms']: (res) => isExists(res, "data[].symptoms"),
        ['should all have a medications']: (res) => isExists(res, "data[].medications"),
        ['should all have a createdBy']: (res) => isExists(res, "data[].createdBy"),
        ['should all have a createdBy.name']: (res) => isExists(res, "data[].createdBy.name"),
        ['should all have a createdBy.nip']: (res) => isExists(res, "data[].createdBy.nip", userNurse.nip),
        ['should all have a createdBy.userId']: (res) => isExists(res, "data[].createdBy.userId"),
        ['should have createdAt and format should be date']: (res) => isEqualWith(res, 'data[].createdAt', (v) => v.every(a => isValidDate(a))),
        ['should not have more than 5 result']: (res) => isTotalDataInRange(res, 'data[]', 1, 5),
    }, config, tags);

    testGetAssert(currentFeature, "get all record with createdAt asc", currentRoute, { createdAt: 'asc' }, headers, {
        ['should return 200']: (res) => res.status === 200,
        ['should all have a identityDetail']: (res) => isExists(res, "data[].identityDetail"),
        ['should all have a identityDetail.identityNumber']: (res) => isExists(res, "data[].identityDetail.identityNumber"),
        ['should all have a identityDetail.phoneNumber']: (res) => isExists(res, "data[].identityDetail.phoneNumber"),
        ['should all have a identityDetail.name']: (res) => isExists(res, "data[].identityDetail.name"),
        ['should all have a identityDetail.birthDate and the format should be date']: (res) => isEqualWith(res, 'data[].identityDetail.birthDate', (v) => v.every(a => isValidDate(a))),
        ['should all have a identityDetail.gender']: (res) => isExists(res, "data[].identityDetail.gender"),
        ['should all have a identityDetail.identityCardScanImg']: (res) => isExists(res, "data[].identityDetail.identityCardScanImg"),
        ['should all have a symptoms']: (res) => isExists(res, "data[].symptoms"),
        ['should all have a medications']: (res) => isExists(res, "data[].medications"),
        ['should all have a createdBy']: (res) => isExists(res, "data[].createdBy"),
        ['should all have a createdBy.name']: (res) => isExists(res, "data[].createdBy.name"),
        ['should all have a createdBy.nip']: (res) => isExists(res, "data[].createdBy.nip", userNurse.nip),
        ['should all have a createdBy.userId']: (res) => isExists(res, "data[].createdBy.userId"),
        ['should have createdAt and format should be date']: (res) => isEqualWith(res, 'data[].createdAt', (v) => v.every(a => isValidDate(a))),
        ['should not have more than 5 result']: (res) => isTotalDataInRange(res, 'data[]', 1, 5),
        ['should have return ordered by oldest first']: (res) => isOrdered(res, 'data[].createdAt', 'asc', (v) => new Date(v)),
    }, config, tags);

    const paginationRes = testGetAssert(currentFeature, "get all record with limit", currentRoute, { limit: 2 }, headers, {
        ['should return 200']: (res) => res.status === 200,
        ['should all have a identityDetail']: (res) => isExists(res, "data[].identityDetail"),
        ['should all have a identityDetail.identityNumber']: (res) => isExists(res, "data[].identityDetail.identityNumber"),
        ['should all have a identityDetail.phoneNumber']: (res) => isExists(res, "data[].identityDetail.phoneNumber"),
        ['should all have a identityDetail.name']: (res) => isExists(res, "data[].identityDetail.name"),
        ['should all have a identityDetail.birthDate and the format should be date']: (res) => isEqualWith(res, 'data[].identityDetail.birthDate', (v) => v.every(a => isValidDate(a))),
        ['should all have a identityDetail.gender']: (res) => isExists(res, "data[].identityDetail.gender"),
        ['should all have a identityDetail.identityCardScanImg']: (res) => isExists(res, "data[].identityDetail.identityCardScanImg"),
        ['should all have a symptoms']: (res) => isExists(res, "data[].symptoms"),
        ['should all have a medications']: (res) => isExists(res, "data[].medications"),
        ['should all have a createdBy']: (res) => isExists(res, "data[].createdBy"),
        ['should all have a createdBy.name']: (res) => isExists(res, "data[].createdBy.name"),
        ['should all have a createdBy.nip']: (res) => isExists(res, "data[].createdBy.nip"),
        ['should all have a createdBy.userId']: (res) => isExists(res, "data[].createdBy.userId"),
        ['should have createdAt and format should be date']: (res) => isEqualWith(res, 'data[].createdAt', (v) => v.every(a => isValidDate(a))),
        ['should not have more than 2 result']: (res) => isTotalDataInRange(res, 'data[]', 1, 2),
    }, config, tags);

    if (paginationRes.isSuccess) {
        testGetAssert(currentFeature, "get all record with limit and offset", currentRoute, { limit: 2, offset: 2 }, headers, {
            ['should return 200']: (res) => res.status === 200,
            ['should all have a identityDetail']: (res) => isExists(res, "data[].identityDetail"),
            ['should all have a identityDetail.identityNumber']: (res) => isExists(res, "data[].identityDetail.identityNumber"),
            ['should all have a identityDetail.phoneNumber']: (res) => isExists(res, "data[].identityDetail.phoneNumber"),
            ['should all have a identityDetail.name']: (res) => isExists(res, "data[].identityDetail.name"),
            ['should all have a identityDetail.birthDate and the format should be date']: (res) => isEqualWith(res, 'data[].identityDetail.birthDate', (v) => v.every(a => isValidDate(a))),
            ['should all have a identityDetail.gender']: (res) => isExists(res, "data[].identityDetail.gender"),
            ['should all have a identityDetail.identityCardScanImg']: (res) => isExists(res, "data[].identityDetail.identityCardScanImg"),
            ['should all have a symptoms']: (res) => isExists(res, "data[].symptoms"),
            ['should all have a medications']: (res) => isExists(res, "data[].medications"),
            ['should all have a createdBy']: (res) => isExists(res, "data[].createdBy"),
            ['should all have a createdBy.name']: (res) => isExists(res, "data[].createdBy.name"),
            ['should all have a createdBy.nip']: (res) => isExists(res, "data[].createdBy.nip"),
            ['should all have a createdBy.userId']: (res) => isExists(res, "data[].createdBy.userId"),
            ['should have createdAt and format should be date']: (res) => isEqualWith(res, 'data[].createdAt', (v) => v.every(a => isValidDate(a))),
            ['should not have more than 2 result']: (res) => isTotalDataInRange(res, 'data[]', 1, 2),
            ['should have different data from offset 0']: (res) => {
                try {
                    return res.json().data.every(e => {
                        return paginationRes.res.json().data.every(a => a.createdAt !== e.createdAt)
                    })
                } catch (error) {
                    return false
                }
            },
        }, config, tags);
    }
}