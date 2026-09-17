package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == v2 {
		v27 = v2
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
		if v14 == int32(0) {
			v27 = v2
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != int32(7) {
				v27 = v2
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v20 != int32(17) {
					v27 = v2
				} else {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
					v27 = v23 ^ int32(1)
				}
			}
		}
	}
	if v27&int32(1) != 0 {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v31 = F_get_fn_opclass_options(m, v30)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
			v36 = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
			v46 = int32(4)
			v47 = v45 & v46
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
			if v48&v46 != 0 {
				if v47 != 0 {
					v216 = int32(0)
				} else {
					v53 = v36 << (uint(int32(3)) % 32)
					if v36 <= int32(0) {
						v216 = v53
					} else {
						v152 = v38 + int32(8)
						v153 = v53
						v155 = v152
						v156 = int32(0)
						v159 = int32(0)
						for {
							v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
							v165 = int32(1)
							v200 = v156 + v164&v165 + int32(base.Ui32(v164)>>(uint(int32(7))%32)) + int32(base.Ui32(v164)>>(uint(v165)%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(2))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(3))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(4))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(5))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(6))%32))&v165
							v204 = v159 + v165
							if v204 != v36 {
								v155 = v155 + v165
								v156 = v200
								v159 = v204
								continue
							} else {
								break
							}
							break
						}
						v216 = v153 - v200
					}
				}
			} else {
				if v47 != 0 {
					v59 = v36 << (uint(int32(3)) % 32)
					if v36 <= int32(0) {
						v216 = v59
					} else {
						v152 = v37 + int32(8)
						v153 = v59
						v155 = v152
						v156 = int32(0)
						v159 = int32(0)
						for {
							v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
							v165 = int32(1)
							v200 = v156 + v164&v165 + int32(base.Ui32(v164)>>(uint(int32(7))%32)) + int32(base.Ui32(v164)>>(uint(v165)%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(2))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(3))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(4))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(5))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(6))%32))&v165
							v204 = v159 + v165
							if v204 != v36 {
								v155 = v155 + v165
								v156 = v200
								v159 = v204
								continue
							} else {
								break
							}
							break
						}
						v216 = v153 - v200
					}
				} else {
					if v36 <= int32(0) {
						v216 = int32(0)
					} else {
						v67 = int32(8)
						v68 = v38 + v67
						v70 = v37 + v67
						v71 = int32(0)
						v73 = int32(1)
						v75 = v36 << (uint(int32(3)) % 32)
						if v75 <= v73 {
							v78 = v73
						} else {
							v78 = v75
						}
						if v78 != int32(1) {
							v86 = v71
							v87 = v71
							v88 = int32(0)
							for {
								v96 = int32(base.Ui32(v87) >> (uint(int32(3)) % 32))
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v96))))
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v96))))
								v101 = v98 ^ v100
								v103 = v87 & int32(6)
								v104 = int32(1)
								v113 = int32(base.Ui32(v101)>>(uint(v103|v104)%32))&v104 + (int32(base.Ui32(v101)>>(uint(v103)%32))&v104 + v86)
								v114 = int32(2)
								v115 = v87 + v114
								v117 = v88 + v114
								if v117 != v78&int32(2147483640) {
									v86 = v113
									v87 = v115
									v88 = v117
									continue
								} else {
									break
								}
								break
							}
							if v78&int32(1) == int32(0) {
								v143 = v113
							} else {
								v121 = v113
								v122 = v115
								v131 = int32(base.Ui32(v122) >> (uint(int32(3)) % 32))
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v131))))
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v70))))
								v143 = v121 + int32(base.Ui32(v133^v135)>>(uint(v122&int32(7))%32))&int32(1)
							}
						} else {
							v121 = v71
							v122 = v71
							v131 = int32(base.Ui32(v122) >> (uint(int32(3)) % 32))
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v131))))
							v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v70))))
							v143 = v121 + int32(base.Ui32(v133^v135)>>(uint(v122&int32(7))%32))&int32(1)
						}
						v216 = v143
					}
				}
			}
			*(*float32)(unsafe.Add(mBase, uint32(v6))) = base.F32_convert_i32_s(v216)
			return v6
		}
	} else {
		v36 = int32(16)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
		v46 = int32(4)
		v47 = v45 & v46
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+4)))
		if v48&v46 != 0 {
			if v47 != 0 {
				v216 = int32(0)
			} else {
				v53 = v36 << (uint(int32(3)) % 32)
				if v36 <= int32(0) {
					v216 = v53
				} else {
					v152 = v38 + int32(8)
					v153 = v53
					v155 = v152
					v156 = int32(0)
					v159 = int32(0)
					for {
						v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
						v165 = int32(1)
						v200 = v156 + v164&v165 + int32(base.Ui32(v164)>>(uint(int32(7))%32)) + int32(base.Ui32(v164)>>(uint(v165)%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(2))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(3))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(4))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(5))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(6))%32))&v165
						v204 = v159 + v165
						if v204 != v36 {
							v155 = v155 + v165
							v156 = v200
							v159 = v204
							continue
						} else {
							break
						}
						break
					}
					v216 = v153 - v200
				}
			}
		} else {
			if v47 != 0 {
				v59 = v36 << (uint(int32(3)) % 32)
				if v36 <= int32(0) {
					v216 = v59
				} else {
					v152 = v37 + int32(8)
					v153 = v59
					v155 = v152
					v156 = int32(0)
					v159 = int32(0)
					for {
						v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
						v165 = int32(1)
						v200 = v156 + v164&v165 + int32(base.Ui32(v164)>>(uint(int32(7))%32)) + int32(base.Ui32(v164)>>(uint(v165)%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(2))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(3))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(4))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(5))%32))&v165 + int32(base.Ui32(v164)>>(uint(int32(6))%32))&v165
						v204 = v159 + v165
						if v204 != v36 {
							v155 = v155 + v165
							v156 = v200
							v159 = v204
							continue
						} else {
							break
						}
						break
					}
					v216 = v153 - v200
				}
			} else {
				if v36 <= int32(0) {
					v216 = int32(0)
				} else {
					v67 = int32(8)
					v68 = v38 + v67
					v70 = v37 + v67
					v71 = int32(0)
					v73 = int32(1)
					v75 = v36 << (uint(int32(3)) % 32)
					if v75 <= v73 {
						v78 = v73
					} else {
						v78 = v75
					}
					if v78 != int32(1) {
						v86 = v71
						v87 = v71
						v88 = int32(0)
						for {
							v96 = int32(base.Ui32(v87) >> (uint(int32(3)) % 32))
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v96))))
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v96))))
							v101 = v98 ^ v100
							v103 = v87 & int32(6)
							v104 = int32(1)
							v113 = int32(base.Ui32(v101)>>(uint(v103|v104)%32))&v104 + (int32(base.Ui32(v101)>>(uint(v103)%32))&v104 + v86)
							v114 = int32(2)
							v115 = v87 + v114
							v117 = v88 + v114
							if v117 != v78&int32(2147483640) {
								v86 = v113
								v87 = v115
								v88 = v117
								continue
							} else {
								break
							}
							break
						}
						if v78&int32(1) == int32(0) {
							v143 = v113
						} else {
							v121 = v113
							v122 = v115
							v131 = int32(base.Ui32(v122) >> (uint(int32(3)) % 32))
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v131))))
							v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v70))))
							v143 = v121 + int32(base.Ui32(v133^v135)>>(uint(v122&int32(7))%32))&int32(1)
						}
					} else {
						v121 = v71
						v122 = v71
						v131 = int32(base.Ui32(v122) >> (uint(int32(3)) % 32))
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v131))))
						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v70))))
						v143 = v121 + int32(base.Ui32(v133^v135)>>(uint(v122&int32(7))%32))&int32(1)
					}
					v216 = v143
				}
			}
		}
		*(*float32)(unsafe.Add(mBase, uint32(v6))) = base.F32_convert_i32_s(v216)
		return v6
	}
}
