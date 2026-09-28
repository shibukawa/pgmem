package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui32(v17-int32(2147483641)) < base.Ui32(int32(-2147483640)) {
			v89 = v13
			m.G0 = v10 + int32(16)
			return base.I64_extend_i32_u(v89)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v22 == v17 {
				v89 = v13
				m.G0 = v10 + int32(16)
				return base.I64_extend_i32_u(v89)
			} else {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
				if v24 == int64(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v28 = F_errsave_start(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						if v28 == int32(0) {
							v89 = int32(0)
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v89)
						} else {
							F_errcode(m, int32(101187714))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v36
								F_errmsg(m, int32(_a_F_bit_0), v10)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v27, int32(_a_F_bit_1), int32(407), int32(_a_F_bit_2))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int64(0)
									} else {
										v89 = int32(0)
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v89)
									}
								}
							}
						}
					}
				} else {
					v51 = int32(base.Ui32(v17+int32(7)) >> (uint(int32(3)) % 32))
					v53 = v51 + int32(8)
					v54 = F_palloc0(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v17
						v57 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = v53 << (uint(v57) % 32)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v64 = int32(base.Ui32(v60)>>(uint(v57)%32)) - int32(8)
						if base.Ui32(v51) < base.Ui32(v64) {
							v66 = v51
						} else {
							v66 = v64
						}
						if v66 != 0 {
							v67 = int32(8)
							base.MemoryCopy(m, v54+v67, v13+v67, v66)
						} else {
						}
						v76 = v53<<(uint(int32(3))%32) - v17 + int32(-64)
						if int32(0) < v76 {
							v81 = v54 + v53 - int32(1)
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
							v85 = v82 & (int32(255) << (uint(v76) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v85)
						} else {
						}
						v89 = v54
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v89)
					}
				}
			}
		}
	}
}
func F_bit_bit_count(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v158 int64
	_ = v158
	v6 = int64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(8)
		v13 = v8 + v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v18 = v16 - v12
		if base.Ui32(v14) <= base.Ui32(int32(63)) {
			if v18 == int32(0) {
				return int64(0)
			} else {
				v25 = int32(3)
				v26 = v16 & v25
				if base.Ui32(v25) <= base.Ui32(v16-int32(9)) {
					v34 = v13
					v36 = int32(0)
					v39 = v6
					for {
						v40 = int32(4)
						v41 = v34 + v40
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+3)))
						v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_bit_bit_count[0]))))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+2)))
						v45 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_bit_bit_count[0]))))
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
						v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_bit_bit_count[0]))))
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
						v49 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_bit_bit_count[0]))))
						v53 = v43 + (v45 + (v47 + (v39 + v49)))
						v55 = v36 + v40
						if v55 != v18&int32(-4) {
							v34 = v41
							v36 = v55
							v39 = v53
							continue
						} else {
							break
						}
						break
					}
					if v26 == int32(0) {
						v85 = v53
					} else {
						v59 = v41
						v64 = v53
						v66 = v59
						v67 = int32(0)
						v71 = v64
						for {
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
							v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bit_bit_count[0]))))
							v74 = v71 + v73
							v75 = int32(1)
							v78 = v67 + v75
							if v78 != v26 {
								v66 = v66 + v75
								v67 = v78
								v71 = v74
								continue
							} else {
								break
							}
							break
						}
						v85 = v74
					}
				} else {
					v59 = v13
					v64 = v6
					v66 = v59
					v67 = int32(0)
					v71 = v64
					for {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
						v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bit_bit_count[0]))))
						v74 = v71 + v73
						v75 = int32(1)
						v78 = v67 + v75
						if v78 != v26 {
							v66 = v66 + v75
							v67 = v78
							v71 = v74
							continue
						} else {
							break
						}
						break
					}
					v85 = v74
				}
				return v85
			}
		} else {
			v87 = int64(0)
			v88 = int32(0)
			if v18 == v88 {
				v158 = int64(0)
			} else {
				v95 = v18 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v18) {
					v100 = v13
					v102 = v87
					v105 = v88
					for {
						v106 = int32(4)
						v107 = v100 + v106
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
						v109 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v108)+uint32(_c_F_bit_bit_count[0]))))
						v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
						v111 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_bit_bit_count[0]))))
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
						v113 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v112)+uint32(_c_F_bit_bit_count[0]))))
						v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
						v115 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_bit_bit_count[0]))))
						v119 = v109 + (v111 + (v113 + (v102 + v115)))
						v121 = v105 + v106
						if v121 != v18&int32(-4) {
							v100 = v107
							v102 = v119
							v105 = v121
							continue
						} else {
							break
						}
						break
					}
					if v95 == int32(0) {
						v148 = v119
					} else {
						v125 = v107
						v127 = v119
						v132 = v125
						v133 = int32(0)
						v134 = v127
						for {
							v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
							v139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_bit_bit_count[0]))))
							v140 = v134 + v139
							v141 = int32(1)
							v144 = v133 + v141
							if v144 != v95 {
								v132 = v132 + v141
								v133 = v144
								v134 = v140
								continue
							} else {
								break
							}
							break
						}
						v148 = v140
					}
				} else {
					v125 = v13
					v127 = v87
					v132 = v125
					v133 = int32(0)
					v134 = v127
					for {
						v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
						v139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_bit_bit_count[0]))))
						v140 = v134 + v139
						v141 = int32(1)
						v144 = v133 + v141
						if v144 != v95 {
							v132 = v132 + v141
							v133 = v144
							v134 = v140
							continue
						} else {
							break
						}
						break
					}
					v148 = v140
				}
				v158 = v148
			}
			return v158
		}
	}
}
func F_bit_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v244 int64
	_ = v244
	v2 = int32(0)
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	switch v18 - int32(88) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v27 = v17
		goto L4
	case 10:
		goto L5
	default:
		goto L6
	}
L1:
	;
	m.G0 = v13 - int32(-64)
	return v244
L2:
	;
	if int32(0) < v16 {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	v31 = v17 + int32(1)
	v32 = F_strlen(m, v31)
	mBase = m.M
	if int32(536870911) <= v32 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v29 = F_strlen(m, v27)
	mBase = m.M
	v58 = v27
	v60 = int32(1)
	v61 = v29
	goto L2
L5:
	;
	v27 = v17 + int32(1)
	goto L4
L6:
	;
	if v18 == int32(120) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v18 != int32(66) {
		v27 = v17
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v35 = F_errsave_start(m, v15)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v58 = v31
	v60 = v2
	v61 = v32 << (uint(int32(2)) % 32)
	goto L2
L12:
	;
	return int64(0)
L13:
	;
	if v35 == int32(0) {
		v244 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(2147483640)
	F_errmsg(m, int32(_a_F_bit_in_0), v11+int32(-16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, v15, int32(_a_F_bit_in_1), int32(199), int32(_a_F_bit_in_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v244 = v10
	goto L1
L18:
	;
	v244 = base.I64_extend_i32_u(v72)
	goto L1
L19:
	;
	v176 = v58
	v177 = v79
	v180 = int32(128)
	goto L53
L20:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v102 == int32(0) {
		goto L18
	} else {
		goto L33
	}
L21:
	;
	v83 = F_errsave_start(m, v15)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L28
	}
L22:
	;
	if v61 != v16 {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v65 = v61
	goto L24
L24:
	;
	v68 = int32(8)
	v69 = base.I32_div_s(v65+int32(7), v68)
	v71 = v69 + v68
	v72 = F_palloc0(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L26
	}
L25:
	;
	v65 = v16
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v71 << (uint(int32(2)) % 32)
	v79 = v72 + int32(8)
	if v60 == int32(0) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	if v83 == int32(0) {
		v244 = v10
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(101187714))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v61
	F_errmsg(m, int32(_a_F_bit_in_3), v11+int32(-32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, v15, int32(_a_F_bit_in_1), int32(213), int32(_a_F_bit_in_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	v244 = v10
	goto L1
L33:
	;
	v105 = v58
	v106 = v79
	v110 = v102
	v113 = v2
	goto L34
L34:
	;
	v116 = v110 - int32(48)
	if base.Ui32(v116&int32(255)) < base.Ui32(int32(10)) {
		v137 = v116
		goto L39
	} else {
		goto L40
	}
L35:
	;
	goto L18
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v169)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	if v173 != 0 {
		v105 = v105 + int32(1)
		v106 = v171
		v110 = v173
		v113 = v170
		goto L34
	} else {
		goto L52
	}
L37:
	;
	v169 = v137 << (uint(int32(4)) % 32)
	v170 = int32(1)
	v171 = v106
	goto L36
L38:
	;
	v145 = F_errsave_start(m, v15)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L46
	}
L39:
	;
	if v113 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L40:
	;
	if base.Ui32((v110-int32(65))&int32(255)) <= base.Ui32(int32(5)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = v110 - int32(55)
	goto L39
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(int32(5)) < base.Ui32((v110-int32(97))&int32(255)) {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v137 = v110 - int32(87)
	goto L39
L45:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v169 = v140 | v137
	v170 = int32(0)
	v171 = v106 + int32(1)
	goto L36
L46:
	;
	if v145 == int32(0) {
		v244 = v10
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v152 = F_pg_mblen_cstr(m, v105)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v152
	F_errmsg(m, int32(_a_F_bit_in_4), v11+int32(-48))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, v15, int32(_a_F_bit_in_1), int32(260), int32(_a_F_bit_in_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v244 = v10
	goto L1
L52:
	;
	goto L35
L53:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	switch v186 - int32(48) {
	case 0:
		goto L56
	case 1:
		goto L57
	default:
		goto L55
	}
L54:
	;
	if v186 == int32(0) {
		goto L18
	} else {
		goto L61
	}
L55:
	;
	goto L54
L56:
	;
	v192 = int32(1)
	v197 = int32(base.Ui32(v180)>>(uint(v192)%32)) & int32(127)
	if v197 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v190 = v189 | v180
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v190)
	goto L56
L58:
	;
	v199 = v197
	goto L60
L59:
	;
	v199 = int32(-128)
	goto L60
L60:
	;
	v176 = v176 + v192
	v177 = v177 + base.B2i32(v197 == int32(0))
	v180 = v199
	goto L53
L61:
	;
	v205 = F_errsave_start(m, v15)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	if v205 == int32(0) {
		v244 = v10
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	v212 = F_pg_mblen_cstr(m, v176)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v212
	F_errmsg(m, int32(_a_F_bit_in_5), v13)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	F_errsave_finish(m, v15, int32(_a_F_bit_in_1), int32(235), int32(_a_F_bit_in_2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	v244 = v10
	goto L1
}
