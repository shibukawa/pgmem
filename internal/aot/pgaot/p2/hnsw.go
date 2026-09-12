package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswAddHeapTid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	v6 = v4 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v6)
	v10 = l0 + v4*int32(6)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+8)) = uint16(v11)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v13
	return
}
func F_HnswLoadElementFromTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v5 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+65)) = uint8(v6)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+67)) = uint8(v10)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+66)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+64)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v12 | v13<<(uint(int32(16))%32)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+68)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v5)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+82)) = uint16(v18)
	if l2 == v5 {
	} else {
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
		if v24 == int32(0) {
		} else {
			v27 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v27)
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v29)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)))
			if v33 == int32(0) {
			} else {
				v36 = int32(2)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v36)
				v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)))
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v38)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+10))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = v40
				v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
				if v42 == int32(0) {
				} else {
					v45 = int32(3)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v45)
					v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)) = uint16(v47)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v49
					v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
					if v51 == int32(0) {
					} else {
						v54 = int32(4)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v54)
						v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
						*(*uint16)(unsafe.Add(mBase, uint32(l0)+26)) = uint16(v56)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+22))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+22)) = v58
						v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)))
						if v60 == int32(0) {
						} else {
							v63 = int32(5)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v63)
							v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)))
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v65)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v67
							v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)))
							if v69 == int32(0) {
							} else {
								v72 = int32(6)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v72)
								v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)))
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+38)) = uint16(v74)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+34))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+34)) = v76
								v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)))
								if v78 == int32(0) {
								} else {
									v81 = int32(7)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v81)
									v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)))
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+44)) = uint16(v83)
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v85
									v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)))
									if v87 == int32(0) {
									} else {
										v90 = int32(8)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v90)
										v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)))
										*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)) = uint16(v92)
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+46))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+46)) = v94
										v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)))
										if v96 == int32(0) {
										} else {
											v99 = int32(9)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v99)
											v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)))
											*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v101)
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v103
											v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)))
											if v105 == int32(0) {
											} else {
												v108 = int32(10)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v108)
												v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+58))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+58)) = v110
												v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)))
												*(*uint16)(unsafe.Add(mBase, uint32(l0)+62)) = uint16(v112)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if l3 != 0 {
		v118 = F_datumCopy(m, l1+int32(72), int32(0), int32(-1))
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v118
			return
		}
	} else {
		return
	}
}
func F_HnswLoadNeighborTids(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v7 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v10 = F_ReadBuffer(m, l2, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_LockBuffer(m, v10, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v10 < int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+(v10^int32(-1))<<(uint(int32(2))%32))))
				v34 = v26
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v34 = v28 + v10<<(uint(int32(13))%32) + int32(-8192)
			}
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+82)))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35<<(uint(int32(2))%32))+20))
			v42 = v39&int32(32767) + v34
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+67)))
			if v43 != v44 {
				v67 = v7
				F_UnlockReleaseBuffer(m, v10)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v67
				}
			} else {
				v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+2)))
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+65)))
				if v46 != (v47+int32(2))*l3 {
					v67 = v7
					F_UnlockReleaseBuffer(m, v10)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						return v67
					}
				} else {
					v53 = F_mul_size(m, v47-l5, l3)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = int32(6)
						v61 = F_mul_size(m, v55, l4)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							if v61 != 0 {
								v63 = F__emscripten_memcpy_bulkmem(m, l1, v42+v53*v55+int32(4), v61)
								mBase = m.M
							} else {
							}
							v67 = int32(1)
							F_UnlockReleaseBuffer(m, v10)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v67
							}
						}
					}
				}
			}
		}
	}
}
func F_HnswNewBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = int32(0)
	v6 = F_ReadBufferExtended(m, l0, l1, int32(-1), v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_LockBuffer(m, v6, int32(2))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v6
		}
	}
}
