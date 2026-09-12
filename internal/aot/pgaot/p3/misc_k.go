package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KeepFileRestoredFromArchive(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v6 = m.G0
	v8 = v6 - int32(2176)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	v17 = F_pg_snprintf(m, v8+int32(1152), int32(1024), int32(177102), v8+int32(16))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = F___fstatat(m, int32(-100), v8+int32(1152), v8+int32(1056), int32(0))
	mBase = m.M
	goto L4
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L65
	}
L4:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = v8 + int32(32)
	v31 = v8 + int32(1152)
	goto L11
L6:
	;
	goto L7
L7:
	;
	v153 = F_durable_rename(m, l0, v8+int32(1152), int32(21))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L41
	}
L8:
	;
	v149 = F_unlink(m, v8+int32(32))
	mBase = m.M
	if v149 != 0 {
		goto L3
	} else {
		goto L40
	}
L9:
	;
	v144 = F_strlen(m, v133)
	mBase = m.M
	goto L8
L11:
	;
	goto L12
L12:
	;
	v38 = int32(1023)
	if (v29^v31)&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
	goto L9
L14:
	;
	v118 = v113
	v119 = v114
	v120 = v115
	goto L36
L15:
	;
	if v108 == int32(0) {
		v133 = v106
		v134 = v107
		goto L13
	} else {
		goto L35
	}
L16:
	;
	v106 = v31
	v107 = v29
	v108 = v38
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v31&int32(3) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v75 == int32(0) {
		v133 = v72
		v134 = v73
		goto L13
	} else {
		goto L28
	}
L20:
	;
	v72 = v31
	v73 = v29
	v74 = v38
	v75 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v51 = v31
	v52 = v29
	v53 = v38
	goto L23
L23:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v55)
	if v55 == int32(0) {
		v113 = v51
		v114 = v52
		v115 = v53
		goto L14
	} else {
		goto L25
	}
L24:
	;
	v72 = v66
	v73 = v60
	v74 = v62
	v75 = v64
	goto L19
L25:
	;
	v59 = int32(1)
	v60 = v52 + v59
	v62 = v53 - v59
	v63 = int32(0)
	v64 = base.B2i32(v62 != v63)
	v66 = v51 + v59
	if v66&int32(3) == v63 {
		v72 = v66
		v73 = v60
		v74 = v62
		v75 = v64
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if v62 != 0 {
		v51 = v66
		v52 = v60
		v53 = v62
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v78 == int32(0) {
		v106 = v72
		v107 = v73
		v108 = v74
		goto L15
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v74) < base.Ui32(int32(4)) {
		v106 = v72
		v107 = v73
		v108 = v74
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v84 = v72
	v85 = v73
	v86 = v74
	goto L31
L31:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v92 = int32(-2139062144)
	if (int32(16843008)-v89|v89)&v92 != v92 {
		v113 = v84
		v114 = v85
		v115 = v86
		goto L14
	} else {
		goto L33
	}
L32:
	;
	v106 = v100
	v107 = v98
	v108 = v102
	goto L15
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v89
	v97 = int32(4)
	v98 = v85 + v97
	v100 = v84 + v97
	v102 = v86 - v97
	if base.Ui32(int32(3)) < base.Ui32(v102) {
		v84 = v100
		v85 = v98
		v86 = v102
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v113 = v106
	v114 = v107
	v115 = v108
	goto L14
L36:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v122)
	if v122 == int32(0) {
		v133 = v118
		v134 = v119
		goto L13
	} else {
		goto L38
	}
L37:
	;
	v133 = v129
	v134 = v127
	goto L13
L38:
	;
	v126 = int32(1)
	v127 = v119 + v126
	v129 = v118 + v126
	v131 = v120 - v126
	if v131 != 0 {
		v118 = v129
		v119 = v127
		v120 = v131
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L7
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	if v156 != int32(2) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v25 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	F_XLogArchiveForceDone(m, l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_XLogArchiveNotify(m, l1)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L42
L47:
	;
	goto L42
L48:
	;
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	if v165 < v167 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_WalSndWakeup(m, int32(1), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L64
	}
L51:
	;
	v173 = v165
	goto L54
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v179 = v176 + v173*int32(96)
	v181 = v179 + int32(164)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = int32(1)
	if v182 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	F_s_lock(m, v181, int32(495659), int32(3596), int32(463716))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v191 = v179 + int32(88)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v192 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+16)) = uint8(v193)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = int32(0)
	v198 = v173 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	if v198 < v200 {
		v173 = v198
		goto L54
	} else {
		goto L63
	}
L63:
	;
	goto L55
L64:
	;
	m.G0 = v8 + int32(2176)
	return
L65:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(1152)
	F_errmsg(m, int32(299321), v8)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(498570), int32(400), int32(344108))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_KnownAssignedXidsCompress(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	v13 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v16 = v14 - v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v16 == v17 {
		if l0 == int32(0) {
			if l1 == int32(0) {
				v68 = *(*int32)(unsafe.Add(mBase, _consts[44]))
				v72 = F_LWLockAcquire(m, v68+int32(512), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					v74 = int32(0)
					if v14 <= v15 {
						v172 = v74
					} else {
						v76 = int32(1)
						v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
						v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
						if v15+v76 != v14 {
							v88 = v74
							v90 = v15
							v93 = int32(0)
							for {
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
								if v100 == int32(1) {
									v103 = int32(2)
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
									v112 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
									v116 = v88 + v112
								} else {
									v116 = v88
								}
								v117 = int32(1)
								v118 = v90 + v117
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
								if v120 == v117 {
									v123 = int32(2)
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
									v132 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
									v136 = v116 + v132
								} else {
									v136 = v116
								}
								v137 = int32(2)
								v138 = v90 + v137
								v140 = v93 + v137
								if v140 != v16&int32(-2) {
									v88 = v136
									v90 = v138
									v93 = v140
									continue
								} else {
									break
								}
								break
							}
							v142 = v136
							v144 = v138
						} else {
							v142 = v74
							v144 = v15
						}
						if v16&v76 == int32(0) {
							v172 = v142
						} else {
							v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
							if v156 != int32(1) {
								v172 = v142
							} else {
								v159 = int32(2)
								v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
								v168 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
								v172 = v142 + v168
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
					v184 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
					if l1 == v184 {
						v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
						F_LWLockRelease(m, v189+int32(512))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							v198 = m.G0
							v199 = int32(16)
							v200 = v198 - v199
							m.G0 = v200
							F___gettimeofday(m, v200)
							mBase = m.M
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
							v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
							m.G0 = v200 + v199
							*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
							return
						}
					} else {
						v198 = m.G0
						v199 = int32(16)
						v200 = v198 - v199
						m.G0 = v200
						F___gettimeofday(m, v200)
						mBase = m.M
						v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
						v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
						m.G0 = v200 + v199
						*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
						return
					}
				}
			} else {
				v74 = int32(0)
				if v14 <= v15 {
					v172 = v74
				} else {
					v76 = int32(1)
					v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
					v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
					if v15+v76 != v14 {
						v88 = v74
						v90 = v15
						v93 = int32(0)
						for {
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
							if v100 == int32(1) {
								v103 = int32(2)
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
								v112 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
								v116 = v88 + v112
							} else {
								v116 = v88
							}
							v117 = int32(1)
							v118 = v90 + v117
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
							if v120 == v117 {
								v123 = int32(2)
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
								v132 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
								v136 = v116 + v132
							} else {
								v136 = v116
							}
							v137 = int32(2)
							v138 = v90 + v137
							v140 = v93 + v137
							if v140 != v16&int32(-2) {
								v88 = v136
								v90 = v138
								v93 = v140
								continue
							} else {
								break
							}
							break
						}
						v142 = v136
						v144 = v138
					} else {
						v142 = v74
						v144 = v15
					}
					if v16&v76 == int32(0) {
						v172 = v142
					} else {
						v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
						if v156 != int32(1) {
							v172 = v142
						} else {
							v159 = int32(2)
							v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
							v168 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
							v172 = v142 + v168
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
				v184 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
				if l1 == v184 {
					v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					F_LWLockRelease(m, v189+int32(512))
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						v198 = m.G0
						v199 = int32(16)
						v200 = v198 - v199
						m.G0 = v200
						F___gettimeofday(m, v200)
						mBase = m.M
						v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
						v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
						m.G0 = v200 + v199
						*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
						return
					}
				} else {
					v198 = m.G0
					v199 = int32(16)
					v200 = v198 - v199
					m.G0 = v200
					F___gettimeofday(m, v200)
					mBase = m.M
					v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
					v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
					m.G0 = v200 + v199
					*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
					return
				}
			}
		} else {
			return
		}
	} else {
		if l0 == int32(2) {
			v23 = int32(4432256)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[843]))
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[843])) = v25 + v26
			if v16 < v17<<(uint(v26)%32) {
				return
			} else {
				if v25&int32(127) == int32(0) {
					if l1 == int32(0) {
						v68 = *(*int32)(unsafe.Add(mBase, _consts[44]))
						v72 = F_LWLockAcquire(m, v68+int32(512), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							v74 = int32(0)
							if v14 <= v15 {
								v172 = v74
							} else {
								v76 = int32(1)
								v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
								v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
								if v15+v76 != v14 {
									v88 = v74
									v90 = v15
									v93 = int32(0)
									for {
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
										if v100 == int32(1) {
											v103 = int32(2)
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
											v112 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
											v116 = v88 + v112
										} else {
											v116 = v88
										}
										v117 = int32(1)
										v118 = v90 + v117
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
										if v120 == v117 {
											v123 = int32(2)
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
											v132 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
											v136 = v116 + v132
										} else {
											v136 = v116
										}
										v137 = int32(2)
										v138 = v90 + v137
										v140 = v93 + v137
										if v140 != v16&int32(-2) {
											v88 = v136
											v90 = v138
											v93 = v140
											continue
										} else {
											break
										}
										break
									}
									v142 = v136
									v144 = v138
								} else {
									v142 = v74
									v144 = v15
								}
								if v16&v76 == int32(0) {
									v172 = v142
								} else {
									v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
									if v156 != int32(1) {
										v172 = v142
									} else {
										v159 = int32(2)
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
										v168 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
										v172 = v142 + v168
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
							v184 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
							if l1 == v184 {
								v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								F_LWLockRelease(m, v189+int32(512))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return
								} else {
									v198 = m.G0
									v199 = int32(16)
									v200 = v198 - v199
									m.G0 = v200
									F___gettimeofday(m, v200)
									mBase = m.M
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
									v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
									m.G0 = v200 + v199
									*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
									return
								}
							} else {
								v198 = m.G0
								v199 = int32(16)
								v200 = v198 - v199
								m.G0 = v200
								F___gettimeofday(m, v200)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
								v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
								m.G0 = v200 + v199
								*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
								return
							}
						}
					} else {
						v74 = int32(0)
						if v14 <= v15 {
							v172 = v74
						} else {
							v76 = int32(1)
							v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
							v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
							if v15+v76 != v14 {
								v88 = v74
								v90 = v15
								v93 = int32(0)
								for {
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
									if v100 == int32(1) {
										v103 = int32(2)
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
										v112 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
										v116 = v88 + v112
									} else {
										v116 = v88
									}
									v117 = int32(1)
									v118 = v90 + v117
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
									if v120 == v117 {
										v123 = int32(2)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
										v132 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
										v136 = v116 + v132
									} else {
										v136 = v116
									}
									v137 = int32(2)
									v138 = v90 + v137
									v140 = v93 + v137
									if v140 != v16&int32(-2) {
										v88 = v136
										v90 = v138
										v93 = v140
										continue
									} else {
										break
									}
									break
								}
								v142 = v136
								v144 = v138
							} else {
								v142 = v74
								v144 = v15
							}
							if v16&v76 == int32(0) {
								v172 = v142
							} else {
								v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
								if v156 != int32(1) {
									v172 = v142
								} else {
									v159 = int32(2)
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
									v168 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
									v172 = v142 + v168
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
						v184 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
						if l1 == v184 {
							v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
							F_LWLockRelease(m, v189+int32(512))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								v198 = m.G0
								v199 = int32(16)
								v200 = v198 - v199
								m.G0 = v200
								F___gettimeofday(m, v200)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
								v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
								m.G0 = v200 + v199
								*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
								return
							}
						} else {
							v198 = m.G0
							v199 = int32(16)
							v200 = v198 - v199
							m.G0 = v200
							F___gettimeofday(m, v200)
							mBase = m.M
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
							v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
							m.G0 = v200 + v199
							*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
							return
						}
					}
				} else {
					return
				}
			}
		} else {
			if l0 != int32(3) {
				if l1 == int32(0) {
					v68 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					v72 = F_LWLockAcquire(m, v68+int32(512), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						v74 = int32(0)
						if v14 <= v15 {
							v172 = v74
						} else {
							v76 = int32(1)
							v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
							v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
							if v15+v76 != v14 {
								v88 = v74
								v90 = v15
								v93 = int32(0)
								for {
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
									if v100 == int32(1) {
										v103 = int32(2)
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
										v112 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
										v116 = v88 + v112
									} else {
										v116 = v88
									}
									v117 = int32(1)
									v118 = v90 + v117
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
									if v120 == v117 {
										v123 = int32(2)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
										v132 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
										v136 = v116 + v132
									} else {
										v136 = v116
									}
									v137 = int32(2)
									v138 = v90 + v137
									v140 = v93 + v137
									if v140 != v16&int32(-2) {
										v88 = v136
										v90 = v138
										v93 = v140
										continue
									} else {
										break
									}
									break
								}
								v142 = v136
								v144 = v138
							} else {
								v142 = v74
								v144 = v15
							}
							if v16&v76 == int32(0) {
								v172 = v142
							} else {
								v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
								if v156 != int32(1) {
									v172 = v142
								} else {
									v159 = int32(2)
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
									v168 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
									v172 = v142 + v168
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
						v184 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
						if l1 == v184 {
							v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
							F_LWLockRelease(m, v189+int32(512))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								v198 = m.G0
								v199 = int32(16)
								v200 = v198 - v199
								m.G0 = v200
								F___gettimeofday(m, v200)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
								v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
								m.G0 = v200 + v199
								*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
								return
							}
						} else {
							v198 = m.G0
							v199 = int32(16)
							v200 = v198 - v199
							m.G0 = v200
							F___gettimeofday(m, v200)
							mBase = m.M
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
							v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
							m.G0 = v200 + v199
							*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
							return
						}
					}
				} else {
					v74 = int32(0)
					if v14 <= v15 {
						v172 = v74
					} else {
						v76 = int32(1)
						v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
						v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
						if v15+v76 != v14 {
							v88 = v74
							v90 = v15
							v93 = int32(0)
							for {
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
								if v100 == int32(1) {
									v103 = int32(2)
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
									v112 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
									v116 = v88 + v112
								} else {
									v116 = v88
								}
								v117 = int32(1)
								v118 = v90 + v117
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
								if v120 == v117 {
									v123 = int32(2)
									v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
									v132 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
									v136 = v116 + v132
								} else {
									v136 = v116
								}
								v137 = int32(2)
								v138 = v90 + v137
								v140 = v93 + v137
								if v140 != v16&int32(-2) {
									v88 = v136
									v90 = v138
									v93 = v140
									continue
								} else {
									break
								}
								break
							}
							v142 = v136
							v144 = v138
						} else {
							v142 = v74
							v144 = v15
						}
						if v16&v76 == int32(0) {
							v172 = v142
						} else {
							v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
							if v156 != int32(1) {
								v172 = v142
							} else {
								v159 = int32(2)
								v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
								v168 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
								v172 = v142 + v168
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
					v184 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
					if l1 == v184 {
						v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
						F_LWLockRelease(m, v189+int32(512))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							v198 = m.G0
							v199 = int32(16)
							v200 = v198 - v199
							m.G0 = v200
							F___gettimeofday(m, v200)
							mBase = m.M
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
							v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
							m.G0 = v200 + v199
							*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
							return
						}
					} else {
						v198 = m.G0
						v199 = int32(16)
						v200 = v198 - v199
						m.G0 = v200
						F___gettimeofday(m, v200)
						mBase = m.M
						v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
						v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
						m.G0 = v200 + v199
						*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
						return
					}
				}
			} else {
				v39 = *(*int64)(unsafe.Add(mBase, _consts[842]))
				if v39 == int64(0) {
					if l1 == int32(0) {
						v68 = *(*int32)(unsafe.Add(mBase, _consts[44]))
						v72 = F_LWLockAcquire(m, v68+int32(512), int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return
						} else {
							v74 = int32(0)
							if v14 <= v15 {
								v172 = v74
							} else {
								v76 = int32(1)
								v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
								v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
								if v15+v76 != v14 {
									v88 = v74
									v90 = v15
									v93 = int32(0)
									for {
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
										if v100 == int32(1) {
											v103 = int32(2)
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
											v112 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
											v116 = v88 + v112
										} else {
											v116 = v88
										}
										v117 = int32(1)
										v118 = v90 + v117
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
										if v120 == v117 {
											v123 = int32(2)
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
											v132 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
											v136 = v116 + v132
										} else {
											v136 = v116
										}
										v137 = int32(2)
										v138 = v90 + v137
										v140 = v93 + v137
										if v140 != v16&int32(-2) {
											v88 = v136
											v90 = v138
											v93 = v140
											continue
										} else {
											break
										}
										break
									}
									v142 = v136
									v144 = v138
								} else {
									v142 = v74
									v144 = v15
								}
								if v16&v76 == int32(0) {
									v172 = v142
								} else {
									v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
									if v156 != int32(1) {
										v172 = v142
									} else {
										v159 = int32(2)
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
										v168 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
										v172 = v142 + v168
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
							v184 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
							if l1 == v184 {
								v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								F_LWLockRelease(m, v189+int32(512))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return
								} else {
									v198 = m.G0
									v199 = int32(16)
									v200 = v198 - v199
									m.G0 = v200
									F___gettimeofday(m, v200)
									mBase = m.M
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
									v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
									m.G0 = v200 + v199
									*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
									return
								}
							} else {
								v198 = m.G0
								v199 = int32(16)
								v200 = v198 - v199
								m.G0 = v200
								F___gettimeofday(m, v200)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
								v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
								m.G0 = v200 + v199
								*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
								return
							}
						}
					} else {
						v74 = int32(0)
						if v14 <= v15 {
							v172 = v74
						} else {
							v76 = int32(1)
							v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
							v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
							if v15+v76 != v14 {
								v88 = v74
								v90 = v15
								v93 = int32(0)
								for {
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
									if v100 == int32(1) {
										v103 = int32(2)
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
										v112 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
										v116 = v88 + v112
									} else {
										v116 = v88
									}
									v117 = int32(1)
									v118 = v90 + v117
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
									if v120 == v117 {
										v123 = int32(2)
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
										v132 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
										v136 = v116 + v132
									} else {
										v136 = v116
									}
									v137 = int32(2)
									v138 = v90 + v137
									v140 = v93 + v137
									if v140 != v16&int32(-2) {
										v88 = v136
										v90 = v138
										v93 = v140
										continue
									} else {
										break
									}
									break
								}
								v142 = v136
								v144 = v138
							} else {
								v142 = v74
								v144 = v15
							}
							if v16&v76 == int32(0) {
								v172 = v142
							} else {
								v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
								if v156 != int32(1) {
									v172 = v142
								} else {
									v159 = int32(2)
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
									v168 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
									v172 = v142 + v168
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
						v184 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
						if l1 == v184 {
							v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
							F_LWLockRelease(m, v189+int32(512))
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return
							} else {
								v198 = m.G0
								v199 = int32(16)
								v200 = v198 - v199
								m.G0 = v200
								F___gettimeofday(m, v200)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
								v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
								m.G0 = v200 + v199
								*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
								return
							}
						} else {
							v198 = m.G0
							v199 = int32(16)
							v200 = v198 - v199
							m.G0 = v200
							F___gettimeofday(m, v200)
							mBase = m.M
							v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
							v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
							m.G0 = v200 + v199
							*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
							return
						}
					}
				} else {
					v45 = m.G0
					v46 = int32(16)
					v47 = v45 - v46
					m.G0 = v47
					F___gettimeofday(m, v47)
					mBase = m.M
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					v51 = int64(*(*int32)(unsafe.Add(mBase, uint32(v47)+8)))
					m.G0 = v47 + v46
					if v51+v50*int64(1000000)-int64(946684800000000) < v39+int64(1000000) {
						return
					} else {
						if l1 == int32(0) {
							v68 = *(*int32)(unsafe.Add(mBase, _consts[44]))
							v72 = F_LWLockAcquire(m, v68+int32(512), int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								v74 = int32(0)
								if v14 <= v15 {
									v172 = v74
								} else {
									v76 = int32(1)
									v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
									v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
									if v15+v76 != v14 {
										v88 = v74
										v90 = v15
										v93 = int32(0)
										for {
											v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
											if v100 == int32(1) {
												v103 = int32(2)
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
												*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
												v112 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
												v116 = v88 + v112
											} else {
												v116 = v88
											}
											v117 = int32(1)
											v118 = v90 + v117
											v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
											if v120 == v117 {
												v123 = int32(2)
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
												*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
												v132 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
												v136 = v116 + v132
											} else {
												v136 = v116
											}
											v137 = int32(2)
											v138 = v90 + v137
											v140 = v93 + v137
											if v140 != v16&int32(-2) {
												v88 = v136
												v90 = v138
												v93 = v140
												continue
											} else {
												break
											}
											break
										}
										v142 = v136
										v144 = v138
									} else {
										v142 = v74
										v144 = v15
									}
									if v16&v76 == int32(0) {
										v172 = v142
									} else {
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
										if v156 != int32(1) {
											v172 = v142
										} else {
											v159 = int32(2)
											v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
											v168 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
											v172 = v142 + v168
										}
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
								v184 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
								if l1 == v184 {
									v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
									F_LWLockRelease(m, v189+int32(512))
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return
									} else {
										v198 = m.G0
										v199 = int32(16)
										v200 = v198 - v199
										m.G0 = v200
										F___gettimeofday(m, v200)
										mBase = m.M
										v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
										v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
										m.G0 = v200 + v199
										*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
										return
									}
								} else {
									v198 = m.G0
									v199 = int32(16)
									v200 = v198 - v199
									m.G0 = v200
									F___gettimeofday(m, v200)
									mBase = m.M
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
									v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
									m.G0 = v200 + v199
									*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
									return
								}
							}
						} else {
							v74 = int32(0)
							if v14 <= v15 {
								v172 = v74
							} else {
								v76 = int32(1)
								v79 = *(*int32)(unsafe.Add(mBase, _consts[840]))
								v81 = *(*int32)(unsafe.Add(mBase, _consts[841]))
								if v15+v76 != v14 {
									v88 = v74
									v90 = v15
									v93 = int32(0)
									for {
										v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90+v81))))
										if v100 == int32(1) {
											v103 = int32(2)
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v79+v90<<(uint(v103)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v88<<(uint(v103)%32)))) = v109
											v112 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v112)
											v116 = v88 + v112
										} else {
											v116 = v88
										}
										v117 = int32(1)
										v118 = v90 + v117
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v118))))
										if v120 == v117 {
											v123 = int32(2)
											v129 = *(*int32)(unsafe.Add(mBase, uint32(v79+v118<<(uint(v123)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v79+v116<<(uint(v123)%32)))) = v129
											v132 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v116+v81))) = uint8(v132)
											v136 = v116 + v132
										} else {
											v136 = v116
										}
										v137 = int32(2)
										v138 = v90 + v137
										v140 = v93 + v137
										if v140 != v16&int32(-2) {
											v88 = v136
											v90 = v138
											v93 = v140
											continue
										} else {
											break
										}
										break
									}
									v142 = v136
									v144 = v138
								} else {
									v142 = v74
									v144 = v15
								}
								if v16&v76 == int32(0) {
									v172 = v142
								} else {
									v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v81))))
									if v156 != int32(1) {
										v172 = v142
									} else {
										v159 = int32(2)
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v79+v144<<(uint(v159)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v79+v142<<(uint(v159)%32)))) = v165
										v168 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v142+v81))) = uint8(v168)
										v172 = v142 + v168
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v172
							v184 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v184
							if l1 == v184 {
								v189 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								F_LWLockRelease(m, v189+int32(512))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return
								} else {
									v198 = m.G0
									v199 = int32(16)
									v200 = v198 - v199
									m.G0 = v200
									F___gettimeofday(m, v200)
									mBase = m.M
									v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
									v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
									m.G0 = v200 + v199
									*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
									return
								}
							} else {
								v198 = m.G0
								v199 = int32(16)
								v200 = v198 - v199
								m.G0 = v200
								F___gettimeofday(m, v200)
								mBase = m.M
								v203 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
								v204 = int64(*(*int32)(unsafe.Add(mBase, uint32(v200)+8)))
								m.G0 = v200 + v199
								*(*int64)(unsafe.Add(mBase, _consts[842])) = v204 + v203*int64(1000000) - int64(946684800000000)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_koi8r_to_iso(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(22), int32(25))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(22), int32(25), int32(2229296), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
