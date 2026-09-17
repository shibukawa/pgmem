package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_del_members(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v148 int32
	_ = v148
	v3 = int32(0)
	if l0 == v3 {
		return int32(0)
	} else {
		if l1 == int32(0) {
			return l0
		} else {
			v19 = int32(8)
			v20 = l0 + v19
			v22 = l1 + v19
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v24 < v23 {
				v28 = v3
				for {
					v39 = v28 << (uint(int32(2)) % 32)
					v40 = v20 + v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v22)))
					*(*int32)(unsafe.Add(mBase, uint32(v40))) = v41 & (v43 ^ int32(-1))
					v49 = v28 + int32(1)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v49 < v50 {
						v28 = v49
						continue
					} else {
						break
					}
					break
				}
				return l0
			} else {
				v52 = int32(-1)
				if int32(2) <= v23 {
					v55 = int32(1)
					if v23 <= v55 {
						v58 = v55
					} else {
						v58 = v23
					}
					v63 = int32(0)
					v66 = v52
					v67 = v63
					v68 = v63
					for {
						v77 = int32(2)
						v78 = v67 << (uint(v77) % 32)
						v79 = v20 + v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v22+v78)))
						v83 = int32(-1)
						v85 = v80 & (v82 ^ v83)
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = v85
						v88 = v67 | int32(1)
						v90 = v88 << (uint(v77) % 32)
						v91 = v20 + v90
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v22+v90)))
						v97 = v92 & (v94 ^ v83)
						*(*int32)(unsafe.Add(mBase, uint32(v91))) = v97
						if v85 != 0 {
							v99 = v67
						} else {
							v99 = v66
						}
						if v97 != 0 {
							v100 = v88
						} else {
							v100 = v99
						}
						v101 = int32(2)
						v102 = v67 + v101
						v104 = v68 + v101
						if v104 != v58&int32(2147483646) {
							v66 = v100
							v67 = v102
							v68 = v104
							continue
						} else {
							break
						}
						break
					}
					if v58&int32(1) == int32(0) {
						v132 = v100
					} else {
						v109 = v100
						v110 = v102
						v121 = v110 << (uint(int32(2)) % 32)
						v122 = v20 + v121
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v22)))
						v128 = v123 & (v125 ^ int32(-1))
						*(*int32)(unsafe.Add(mBase, uint32(v122))) = v128
						if v128 != 0 {
							v130 = v110
						} else {
							v130 = v109
						}
						v132 = v130
					}
				} else {
					v109 = v52
					v110 = v3
					v121 = v110 << (uint(int32(2)) % 32)
					v122 = v20 + v121
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v22)))
					v128 = v123 & (v125 ^ int32(-1))
					*(*int32)(unsafe.Add(mBase, uint32(v122))) = v128
					if v128 != 0 {
						v130 = v110
					} else {
						v130 = v109
					}
					v132 = v130
				}
				if v132 == int32(-1) {
					F_pfree(m, l0)
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v132 + int32(1)
					return l0
				}
			}
		}
	}
}
func F_bms_difference(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v218 int32
	_ = v218
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v218
L5:
	;
	v23 = v17<<(uint(int32(2))%32) + int32(8)
	v24 = F_palloc(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32 < v17 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return int32(0)
L9:
	;
	if v23 == int32(0) {
		v218 = v24
		goto L4
	} else {
		goto L10
	}
L10:
	;
	base.MemoryCopy(m, v24, l0, v23)
	return v24
L11:
	;
	v83 = v17<<(uint(int32(2))%32) + int32(8)
	v84 = F_palloc(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L20
	}
L12:
	;
	v34 = int32(1)
	if v17 <= v34 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = v34
	goto L15
L14:
	;
	v37 = v17
	goto L15
L15:
	;
	v38 = int32(8)
	v44 = v3
	goto L16
L16:
	;
	v55 = v44 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1+v38+v55)))
	if v57&(v59^int32(-1)) != 0 {
		goto L11
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	v64 = v44 + int32(1)
	if v64 != v37 {
		v44 = v64
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v83 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	base.MemoryCopy(m, v84, l0, v83)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v87 = int32(8)
	v88 = l1 + v87
	v90 = v84 + v87
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v92 < v91 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v97 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v121 = int32(0)
	v122 = int32(-1)
	if int32(2) <= v91 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v108 = v97 << (uint(int32(2)) % 32)
	v109 = v90 + v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110 & (v112 ^ int32(-1))
	v118 = v97 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v118 < v119 {
		v97 = v118
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v218 = v84
	goto L4
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v200 + int32(1)
	v218 = v84
	goto L4
L31:
	;
	v125 = int32(1)
	if v91 <= v125 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v177 = v122
	v179 = v121
	goto L33
L33:
	;
	v190 = v179 << (uint(int32(2)) % 32)
	v191 = v90 + v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v88)))
	v197 = v192 & (v194 ^ int32(-1))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v197
	if v197 != 0 {
		goto L47
	} else {
		goto L48
	}
L34:
	;
	v128 = v125
	goto L36
L35:
	;
	v128 = v91
	goto L36
L36:
	;
	v134 = v122
	v135 = int32(0)
	v136 = v121
	goto L37
L37:
	;
	v146 = int32(2)
	v147 = v136 << (uint(v146) % 32)
	v148 = v90 + v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v88+v147)))
	v152 = int32(-1)
	v154 = v149 & (v151 ^ v152)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v154
	v157 = v136 | int32(1)
	v159 = v157 << (uint(v146) % 32)
	v160 = v90 + v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v88+v159)))
	v166 = v161 & (v163 ^ v152)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v166
	if v154 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v128&int32(1) == int32(0) {
		v200 = v169
		goto L30
	} else {
		goto L46
	}
L39:
	;
	v168 = v136
	goto L41
L40:
	;
	v168 = v134
	goto L41
L41:
	;
	if v166 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v169 = v157
	goto L44
L43:
	;
	v169 = v168
	goto L44
L44:
	;
	v170 = int32(2)
	v171 = v136 + v170
	v173 = v135 + v170
	if v173 != v128&int32(2147483646) {
		v134 = v169
		v135 = v173
		v136 = v171
		goto L37
	} else {
		goto L45
	}
L45:
	;
	goto L38
L46:
	;
	v177 = v169
	v179 = v171
	goto L33
L47:
	;
	v199 = v179
	goto L49
L48:
	;
	v199 = v177
	goto L49
L49:
	;
	v200 = v199
	goto L30
}
func F_bms_get_singleton_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= v12 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v12
	goto L6
L5:
	;
	v16 = v13
	goto L6
L6:
	;
	v21 = int32(0)
	v23 = int32(-1)
	goto L8
L7:
	;
	return v50
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v21<<(uint(int32(2))%32))))
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
	v50 = int32(1)
	goto L7
L10:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v31)))|base.B2i32(int32(0) <= v23) != 0 {
		v50 = v3
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v42 = v23
	goto L12
L12:
	;
	v44 = v21 + int32(1)
	if v44 != v16 {
		v21 = v44
		v23 = v42
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v42 = base.I32_ctz(v31) | v21<<(uint(int32(5))%32)
	goto L12
L14:
	;
	goto L9
}
func F_bms_intersect(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) == v3 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v22 = base.B2i32(v21 < v20)
		if v20 < v21 {
			v24 = v20
		} else {
			v24 = v21
		}
		v28 = v24<<(uint(int32(2))%32) + int32(8)
		v29 = F_palloc(m, v28)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			if v28 != 0 {
				if v21 < v20 {
					v33 = l1
				} else {
					v33 = l0
				}
				base.MemoryCopy(m, v29, v33, v28)
			} else {
			}
			if v21 < v20 {
				v35 = l0
			} else {
				v35 = l1
			}
			v36 = int32(8)
			v37 = v35 + v36
			v39 = v29 + v36
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			if v40 < int32(2) {
				v94 = int32(0)
				v95 = int32(-1)
				v107 = v94 << (uint(int32(2)) % 32)
				v108 = v39 + v107
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v37+v107)))
				v112 = v109 & v111
				*(*int32)(unsafe.Add(mBase, uint32(v108))) = v112
				if v112 != 0 {
					v114 = v94
				} else {
					v114 = v95
				}
				v116 = v114
			} else {
				v45 = int32(1)
				if v40 <= v45 {
					v48 = v45
				} else {
					v48 = v40
				}
				v55 = int32(0)
				v56 = int32(-1)
				v60 = v3
				for {
					v67 = int32(2)
					v68 = v55 << (uint(v67) % 32)
					v69 = v39 + v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v37+v68)))
					v73 = v70 & v72
					*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
					v76 = v55 | int32(1)
					v78 = v76 << (uint(v67) % 32)
					v79 = v39 + v78
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v37+v78)))
					v83 = v80 & v82
					*(*int32)(unsafe.Add(mBase, uint32(v79))) = v83
					if v73 != 0 {
						v85 = v55
					} else {
						v85 = v56
					}
					if v83 != 0 {
						v86 = v76
					} else {
						v86 = v85
					}
					v87 = int32(2)
					v88 = v55 + v87
					v90 = v60 + v87
					if v90 != v48&int32(2147483646) {
						v55 = v88
						v56 = v86
						v60 = v90
						continue
					} else {
						break
					}
					break
				}
				if v48&int32(1) == int32(0) {
					v116 = v86
				} else {
					v94 = v88
					v95 = v86
					v107 = v94 << (uint(int32(2)) % 32)
					v108 = v39 + v107
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v37+v107)))
					v112 = v109 & v111
					*(*int32)(unsafe.Add(mBase, uint32(v108))) = v112
					if v112 != 0 {
						v114 = v94
					} else {
						v114 = v95
					}
					v116 = v114
				}
			}
			if v116 == int32(-1) {
				F_pfree(m, v29)
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v116 + int32(1)
				v139 = v29
				return v139
			}
		}
	} else {
		v139 = v3
		return v139
	}
}
func F_bms_member_index(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v7 = int32(-1)
	if int32(0) <= l1 {
		if l0 == int32(0) {
			v62 = v7
		} else {
			v13 = int32(base.Ui32(l1) >> (uint(int32(5)) % 32))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v14 <= v13 {
				v62 = v7
			} else {
				v17 = l0 + int32(8)
				v20 = v17 + v13<<(uint(int32(2))%32)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				if int32(base.Ui32(v21)>>(uint(l1)%32))&int32(1) == int32(0) {
					v62 = v7
				} else {
					v27 = int32(0)
					v28 = int32(-1)
					if v13 != 0 {
						v34 = int32(0)
						v35 = v27
						for {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v17+v34<<(uint(int32(2))%32))))
							if v42 != 0 {
								v45 = v35 + base.I32_popcnt(v42)
							} else {
								v45 = v35
							}
							v47 = v34 + int32(1)
							if v47 != v13 {
								v34 = v47
								v35 = v45
								continue
							} else {
								break
							}
							break
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
						v52 = v45
						v56 = v49
					} else {
						v52 = v27
						v56 = v21
					}
					v62 = base.I32_popcnt((v28<<(uint(l1)%32)^v28)&v56) + v52
				}
			}
		}
		return v62
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_bms_member_index_0), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_bms_member_index_1), int32(519), int32(_a_F_bms_member_index_2))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_bms_union(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	v3 = int32(0)
	if l0 == v3 {
		if l1 != 0 {
			v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v136 = l1
			v137 = v135
			v141 = v137<<(uint(int32(2))%32) + int32(8)
			v142 = F_palloc(m, v141)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				if v141 == int32(0) {
					v150 = v142
				} else {
					base.MemoryCopy(m, v142, v136, v141)
					v150 = v142
				}
				return v150
			}
		} else {
			return int32(0)
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l1 == int32(0) {
			v136 = l0
			v137 = v16
			v141 = v137<<(uint(int32(2))%32) + int32(8)
			v142 = F_palloc(m, v141)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				if v141 == int32(0) {
					v150 = v142
				} else {
					base.MemoryCopy(m, v142, v136, v141)
					v150 = v142
				}
				return v150
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v20 = base.B2i32(v19 < v16)
			if v19 < v16 {
				v21 = v16
			} else {
				v21 = v19
			}
			v25 = v21<<(uint(int32(2))%32) + int32(8)
			v26 = F_palloc(m, v25)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					if v19 < v16 {
						v30 = l0
					} else {
						v30 = l1
					}
					base.MemoryCopy(m, v26, v30, v25)
				} else {
				}
				if v19 < v16 {
					v33 = l1
				} else {
					v33 = l0
				}
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				if v34 <= int32(1) {
					v37 = int32(1)
				} else {
					v37 = v34
				}
				v39 = v37 & int32(3)
				v40 = int32(8)
				v41 = v33 + v40
				v43 = v26 + v40
				v44 = int32(0)
				if int32(4) <= v34 {
					v51 = v44
					v56 = int32(0)
					for {
						v62 = v51 << (uint(int32(2)) % 32)
						v63 = v43 + v62
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v41+v62)))
						*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 | v66
						v69 = int32(4)
						v70 = v62 | v69
						v71 = v43 + v70
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v41+v70)))
						*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 | v74
						v78 = v62 | int32(8)
						v79 = v43 + v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v41+v78)))
						*(*int32)(unsafe.Add(mBase, uint32(v79))) = v80 | v82
						v86 = v62 | int32(12)
						v87 = v43 + v86
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v41+v86)))
						*(*int32)(unsafe.Add(mBase, uint32(v87))) = v88 | v90
						v94 = v51 + v69
						v96 = v56 + v69
						if v96 != v37&int32(2147483644) {
							v51 = v94
							v56 = v96
							continue
						} else {
							break
						}
						break
					}
					if v39 == int32(0) {
						v150 = v26
					} else {
						v101 = v94
						v112 = v101
						v121 = v3
						for {
							v123 = v112 << (uint(int32(2)) % 32)
							v124 = v43 + v123
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v41+v123)))
							*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v127
							v130 = int32(1)
							v133 = v121 + v130
							if v133 != v39 {
								v112 = v112 + v130
								v121 = v133
								continue
							} else {
								break
							}
							break
						}
						v150 = v26
					}
				} else {
					v101 = v44
					v112 = v101
					v121 = v3
					for {
						v123 = v112 << (uint(int32(2)) % 32)
						v124 = v43 + v123
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v41+v123)))
						*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v127
						v130 = int32(1)
						v133 = v121 + v130
						if v133 != v39 {
							v112 = v112 + v130
							v121 = v133
							continue
						} else {
							break
						}
						break
					}
					v150 = v26
				}
				return v150
			}
		}
	}
}
