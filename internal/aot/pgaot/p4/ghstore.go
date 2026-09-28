package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	v2 = int32(0)
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)))
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
			return int64(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
			v36 = v35
			v39 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8))))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v7))))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
			v49 = int32(4)
			v50 = v48 & v49
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
			if v51&v49 != 0 {
				if v50 != 0 {
					v219 = int32(0)
				} else {
					v56 = v36 << (uint(int32(3)) % 32)
					if v36 <= int32(0) {
						v219 = v56
					} else {
						v155 = v41 + int32(8)
						v156 = v56
						v158 = v155
						v159 = int32(0)
						v162 = int32(0)
						for {
							v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v168 = int32(1)
							v203 = v159 + v167&v168 + int32(base.Ui32(v167)>>(uint(int32(7))%32)) + int32(base.Ui32(v167)>>(uint(v168)%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(2))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(3))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(4))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(5))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(6))%32))&v168
							v207 = v162 + v168
							if v207 != v36 {
								v158 = v158 + v168
								v159 = v203
								v162 = v207
								continue
							} else {
								break
							}
							break
						}
						v219 = v156 - v203
					}
				}
			} else {
				if v50 != 0 {
					v62 = v36 << (uint(int32(3)) % 32)
					if v36 <= int32(0) {
						v219 = v62
					} else {
						v155 = v39 + int32(8)
						v156 = v62
						v158 = v155
						v159 = int32(0)
						v162 = int32(0)
						for {
							v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
							v168 = int32(1)
							v203 = v159 + v167&v168 + int32(base.Ui32(v167)>>(uint(int32(7))%32)) + int32(base.Ui32(v167)>>(uint(v168)%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(2))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(3))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(4))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(5))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(6))%32))&v168
							v207 = v162 + v168
							if v207 != v36 {
								v158 = v158 + v168
								v159 = v203
								v162 = v207
								continue
							} else {
								break
							}
							break
						}
						v219 = v156 - v203
					}
				} else {
					if v36 <= int32(0) {
						v219 = int32(0)
					} else {
						v70 = int32(8)
						v71 = v41 + v70
						v73 = v39 + v70
						v74 = int32(0)
						v76 = int32(1)
						v78 = v36 << (uint(int32(3)) % 32)
						if v78 <= v76 {
							v81 = v76
						} else {
							v81 = v78
						}
						if v81 != int32(1) {
							v89 = v74
							v90 = v74
							v91 = int32(0)
							for {
								v99 = int32(base.Ui32(v90) >> (uint(int32(3)) % 32))
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v99))))
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v99))))
								v104 = v101 ^ v103
								v106 = v90 & int32(6)
								v107 = int32(1)
								v116 = int32(base.Ui32(v104)>>(uint(v106|v107)%32))&v107 + (int32(base.Ui32(v104)>>(uint(v106)%32))&v107 + v89)
								v117 = int32(2)
								v118 = v90 + v117
								v120 = v91 + v117
								if v120 != v81&int32(2147483640) {
									v89 = v116
									v90 = v118
									v91 = v120
									continue
								} else {
									break
								}
								break
							}
							if v81&int32(1) == int32(0) {
								v146 = v116
							} else {
								v124 = v116
								v125 = v118
								v134 = int32(base.Ui32(v125) >> (uint(int32(3)) % 32))
								v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v134))))
								v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v73))))
								v146 = v124 + int32(base.Ui32(v136^v138)>>(uint(v125&int32(7))%32))&int32(1)
							}
						} else {
							v124 = v74
							v125 = v74
							v134 = int32(base.Ui32(v125) >> (uint(int32(3)) % 32))
							v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v134))))
							v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v73))))
							v146 = v124 + int32(base.Ui32(v136^v138)>>(uint(v125&int32(7))%32))&int32(1)
						}
						v219 = v146
					}
				}
			}
			*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v6)))) = base.F32_convert_i32_s(v219)
			return v6
		}
	} else {
		v36 = int32(16)
		v39 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v8))))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v7))))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
		v49 = int32(4)
		v50 = v48 & v49
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
		if v51&v49 != 0 {
			if v50 != 0 {
				v219 = int32(0)
			} else {
				v56 = v36 << (uint(int32(3)) % 32)
				if v36 <= int32(0) {
					v219 = v56
				} else {
					v155 = v41 + int32(8)
					v156 = v56
					v158 = v155
					v159 = int32(0)
					v162 = int32(0)
					for {
						v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
						v168 = int32(1)
						v203 = v159 + v167&v168 + int32(base.Ui32(v167)>>(uint(int32(7))%32)) + int32(base.Ui32(v167)>>(uint(v168)%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(2))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(3))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(4))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(5))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(6))%32))&v168
						v207 = v162 + v168
						if v207 != v36 {
							v158 = v158 + v168
							v159 = v203
							v162 = v207
							continue
						} else {
							break
						}
						break
					}
					v219 = v156 - v203
				}
			}
		} else {
			if v50 != 0 {
				v62 = v36 << (uint(int32(3)) % 32)
				if v36 <= int32(0) {
					v219 = v62
				} else {
					v155 = v39 + int32(8)
					v156 = v62
					v158 = v155
					v159 = int32(0)
					v162 = int32(0)
					for {
						v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
						v168 = int32(1)
						v203 = v159 + v167&v168 + int32(base.Ui32(v167)>>(uint(int32(7))%32)) + int32(base.Ui32(v167)>>(uint(v168)%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(2))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(3))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(4))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(5))%32))&v168 + int32(base.Ui32(v167)>>(uint(int32(6))%32))&v168
						v207 = v162 + v168
						if v207 != v36 {
							v158 = v158 + v168
							v159 = v203
							v162 = v207
							continue
						} else {
							break
						}
						break
					}
					v219 = v156 - v203
				}
			} else {
				if v36 <= int32(0) {
					v219 = int32(0)
				} else {
					v70 = int32(8)
					v71 = v41 + v70
					v73 = v39 + v70
					v74 = int32(0)
					v76 = int32(1)
					v78 = v36 << (uint(int32(3)) % 32)
					if v78 <= v76 {
						v81 = v76
					} else {
						v81 = v78
					}
					if v81 != int32(1) {
						v89 = v74
						v90 = v74
						v91 = int32(0)
						for {
							v99 = int32(base.Ui32(v90) >> (uint(int32(3)) % 32))
							v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v99))))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v99))))
							v104 = v101 ^ v103
							v106 = v90 & int32(6)
							v107 = int32(1)
							v116 = int32(base.Ui32(v104)>>(uint(v106|v107)%32))&v107 + (int32(base.Ui32(v104)>>(uint(v106)%32))&v107 + v89)
							v117 = int32(2)
							v118 = v90 + v117
							v120 = v91 + v117
							if v120 != v81&int32(2147483640) {
								v89 = v116
								v90 = v118
								v91 = v120
								continue
							} else {
								break
							}
							break
						}
						if v81&int32(1) == int32(0) {
							v146 = v116
						} else {
							v124 = v116
							v125 = v118
							v134 = int32(base.Ui32(v125) >> (uint(int32(3)) % 32))
							v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v134))))
							v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v73))))
							v146 = v124 + int32(base.Ui32(v136^v138)>>(uint(v125&int32(7))%32))&int32(1)
						}
					} else {
						v124 = v74
						v125 = v74
						v134 = int32(base.Ui32(v125) >> (uint(int32(3)) % 32))
						v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v134))))
						v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v73))))
						v146 = v124 + int32(base.Ui32(v136^v138)>>(uint(v125&int32(7))%32))&int32(1)
					}
					v219 = v146
				}
			}
		}
		*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v6)))) = base.F32_convert_i32_s(v219)
		return v6
	}
}
