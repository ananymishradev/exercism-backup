package annalyn

func CanFastAttack(knightIsAwake bool) bool {
    if knightIsAwake {
        return false
    }
    return true
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    count := 0
    if knightIsAwake {
        count = count + 1
    }
    if archerIsAwake {
        count = count + 1
    }
    if prisonerIsAwake {
        count = count + 1
    }
    if count >= 1 {
        return true
    }
    return false
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if !archerIsAwake && prisonerIsAwake {
        return true
    }
    return false
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    if petDogIsPresent == true {
        if archerIsAwake == false {
            return true
        }
    }
    if petDogIsPresent == false {
        if prisonerIsAwake == true {
            if knightIsAwake == false && archerIsAwake == false {
                return true
            }
        }
    }
    return false
}
