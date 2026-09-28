package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_KeepFileRestoredFromArchive(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
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
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	v7 = m.G0
	v9 = v7 - int32(2176)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	v13 = v9 + int32(1152)
	v18 = F_pg_snprintf(m, v13, int32(1024), int32(_a_F_KeepFileRestoredFromArchive_0), v9+int32(16))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = F___fstatat(m, int32(-100), v13, v9+int32(1056), int32(0))
	mBase = m.M
	goto L4
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L64
	}
L4:
	;
	if v24 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = v9 + int32(32)
	goto L11
L6:
	;
	goto L7
L7:
	;
	v153 = F_durable_rename(m, l0, v9+int32(1152), int32(21))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L40
	}
L8:
	;
	v148 = F_unlink(m, v28)
	mBase = m.M
	if v148 != 0 {
		goto L3
	} else {
		goto L39
	}
L9:
	;
	v145 = F_strlen(m, v134)
	mBase = m.M
	goto L8
L11:
	;
	goto L12
L12:
	;
	v35 = int32(1023)
	if (v28^v13)&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v138)
	goto L9
L14:
	;
	v119 = v114
	v120 = v115
	v121 = v116
	goto L35
L15:
	;
	if v109 == int32(0) {
		v134 = v107
		v135 = v108
		goto L13
	} else {
		goto L34
	}
L16:
	;
	v107 = v13
	v108 = v28
	v109 = v35
	goto L15
L17:
	;
	goto L18
L18:
	;
	v39 = int32(0)
	if base.B2i32(v13&int32(3) == v39)|int32(0) == v39 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v75 == int32(0) {
		v134 = v72
		v135 = v73
		goto L13
	} else {
		goto L28
	}
L20:
	;
	v51 = v13
	v52 = v28
	v53 = v35
	goto L23
L21:
	;
	goto L22
L22:
	;
	v72 = v13
	v73 = v28
	v74 = v35
	v75 = int32(1)
	goto L19
L23:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v55)
	if v55 == int32(0) {
		v114 = v51
		v115 = v52
		v116 = v53
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
	if base.B2i32(v78 == int32(0))|base.B2i32(base.Ui32(v74) < base.Ui32(int32(4))) != 0 {
		v107 = v72
		v108 = v73
		v109 = v74
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v85 = v72
	v86 = v73
	v87 = v74
	goto L30
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 != v93 {
		v114 = v85
		v115 = v86
		v116 = v87
		goto L14
	} else {
		goto L32
	}
L31:
	;
	v107 = v101
	v108 = v99
	v109 = v103
	goto L15
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v90
	v98 = int32(4)
	v99 = v86 + v98
	v101 = v85 + v98
	v103 = v87 - v98
	if base.Ui32(int32(3)) < base.Ui32(v103) {
		v85 = v101
		v86 = v99
		v87 = v103
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v114 = v107
	v115 = v108
	v116 = v109
	goto L14
L35:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v123)
	if v123 == int32(0) {
		v134 = v119
		v135 = v120
		goto L13
	} else {
		goto L37
	}
L36:
	;
	v134 = v130
	v135 = v128
	goto L13
L37:
	;
	v127 = int32(1)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if v132 != 0 {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L7
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_KeepFileRestoredFromArchive[0]))
	if v156 != int32(2) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v24 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	F_XLogArchiveForceDone(m, l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_XLogArchiveNotify(m, l1)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	goto L41
L47:
	;
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_KeepFileRestoredFromArchive[1]))
	if v165 < v167 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	F_WalSndWakeup(m, int32(1), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v171 = v165
	goto L53
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_KeepFileRestoredFromArchive[2]))
	v180 = v177 + v171*int32(96)
	v181 = int32(164)
	v182 = v180 + v181
	v185 = base.AtomicRmwXchg32(m, v180, v181, int32(1))
	if v185 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	F_s_lock(m, v182, int32(_a_F_KeepFileRestoredFromArchive_1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v190 = v180 + int32(88)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	if v191 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+16)) = uint8(v192)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v194 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v182))), uint32(v194))
	v198 = v171 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_KeepFileRestoredFromArchive[1]))
	if v198 < v200 {
		v171 = v198
		goto L53
	} else {
		goto L62
	}
L62:
	;
	goto L54
L63:
	;
	m.G0 = v9 + int32(2176)
	return
L64:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(1152)
	F_errmsg(m, int32(_a_F_KeepFileRestoredFromArchive_2), v9)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_KeepFileRestoredFromArchive_3), int32(401), int32(_a_F_KeepFileRestoredFromArchive_4))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_KnownAssignedXidsCompress(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v13 = v11 - v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v13 == v14 {
		if l0 == int32(0) {
			if l1 == int32(0) {
				v65 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
				v69 = F_LWLockAcquire(m, v65+int32(512), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					v71 = int32(0)
					if v12 < v11 {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
						v75 = v71
						v77 = v74
						v78 = v12
						for {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
							if v84 == int32(1) {
								v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
								v89 = int32(2)
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
								v97 = int32(_a_F_KnownAssignedXidsCompress_0)
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v100 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
								v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v106 = v75 + v100
								v107 = v103
							} else {
								v106 = v75
								v107 = v77
							}
							v109 = v78 + int32(1)
							if v109 != v11 {
								v75 = v106
								v77 = v107
								v78 = v109
								continue
							} else {
								break
							}
							break
						}
						v111 = v106
					} else {
						v111 = v71
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
					v120 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
					if l1 == v120 {
						v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
						F_LWLockRelease(m, v125+int32(512))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							v134 = m.G0
							v135 = int32(16)
							v136 = v134 - v135
							m.G0 = v136
							F_gettimeofday(m, v136)
							mBase = m.M
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
							v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
							m.G0 = v136 + v135
							*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
							return
						}
					} else {
						v134 = m.G0
						v135 = int32(16)
						v136 = v134 - v135
						m.G0 = v136
						F_gettimeofday(m, v136)
						mBase = m.M
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
						v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
						m.G0 = v136 + v135
						*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
						return
					}
				}
			} else {
				v71 = int32(0)
				if v12 < v11 {
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
					v75 = v71
					v77 = v74
					v78 = v12
					for {
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
						if v84 == int32(1) {
							v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
							v89 = int32(2)
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
							v97 = int32(_a_F_KnownAssignedXidsCompress_0)
							v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
							v100 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
							v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
							v106 = v75 + v100
							v107 = v103
						} else {
							v106 = v75
							v107 = v77
						}
						v109 = v78 + int32(1)
						if v109 != v11 {
							v75 = v106
							v77 = v107
							v78 = v109
							continue
						} else {
							break
						}
						break
					}
					v111 = v106
				} else {
					v111 = v71
				}
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
				v120 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
				if l1 == v120 {
					v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
					F_LWLockRelease(m, v125+int32(512))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						v134 = m.G0
						v135 = int32(16)
						v136 = v134 - v135
						m.G0 = v136
						F_gettimeofday(m, v136)
						mBase = m.M
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
						v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
						m.G0 = v136 + v135
						*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
						return
					}
				} else {
					v134 = m.G0
					v135 = int32(16)
					v136 = v134 - v135
					m.G0 = v136
					F_gettimeofday(m, v136)
					mBase = m.M
					v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
					v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
					m.G0 = v136 + v135
					*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
					return
				}
			}
		} else {
			return
		}
	} else {
		if l0 == int32(2) {
			v20 = int32(_a_F_KnownAssignedXidsCompress_1)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[5]))
			v23 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[5])) = v22 + v23
			if v13 < v14<<(uint(v23)%32) {
				return
			} else {
				if v22&int32(127) == int32(0) {
					if l1 == int32(0) {
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
						v69 = F_LWLockAcquire(m, v65+int32(512), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							v71 = int32(0)
							if v12 < v11 {
								v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v75 = v71
								v77 = v74
								v78 = v12
								for {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
									if v84 == int32(1) {
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
										v89 = int32(2)
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
										v97 = int32(_a_F_KnownAssignedXidsCompress_0)
										v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
										v100 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
										v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
										v106 = v75 + v100
										v107 = v103
									} else {
										v106 = v75
										v107 = v77
									}
									v109 = v78 + int32(1)
									if v109 != v11 {
										v75 = v106
										v77 = v107
										v78 = v109
										continue
									} else {
										break
									}
									break
								}
								v111 = v106
							} else {
								v111 = v71
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
							v120 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
							if l1 == v120 {
								v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
								F_LWLockRelease(m, v125+int32(512))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									v134 = m.G0
									v135 = int32(16)
									v136 = v134 - v135
									m.G0 = v136
									F_gettimeofday(m, v136)
									mBase = m.M
									v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
									v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
									m.G0 = v136 + v135
									*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
									return
								}
							} else {
								v134 = m.G0
								v135 = int32(16)
								v136 = v134 - v135
								m.G0 = v136
								F_gettimeofday(m, v136)
								mBase = m.M
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
								v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
								m.G0 = v136 + v135
								*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
								return
							}
						}
					} else {
						v71 = int32(0)
						if v12 < v11 {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
							v75 = v71
							v77 = v74
							v78 = v12
							for {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
								if v84 == int32(1) {
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
									v89 = int32(2)
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
									v97 = int32(_a_F_KnownAssignedXidsCompress_0)
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v100 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
									v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v106 = v75 + v100
									v107 = v103
								} else {
									v106 = v75
									v107 = v77
								}
								v109 = v78 + int32(1)
								if v109 != v11 {
									v75 = v106
									v77 = v107
									v78 = v109
									continue
								} else {
									break
								}
								break
							}
							v111 = v106
						} else {
							v111 = v71
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
						v120 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
						if l1 == v120 {
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
							F_LWLockRelease(m, v125+int32(512))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								v134 = m.G0
								v135 = int32(16)
								v136 = v134 - v135
								m.G0 = v136
								F_gettimeofday(m, v136)
								mBase = m.M
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
								v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
								m.G0 = v136 + v135
								*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
								return
							}
						} else {
							v134 = m.G0
							v135 = int32(16)
							v136 = v134 - v135
							m.G0 = v136
							F_gettimeofday(m, v136)
							mBase = m.M
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
							v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
							m.G0 = v136 + v135
							*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
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
					v65 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
					v69 = F_LWLockAcquire(m, v65+int32(512), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = int32(0)
						if v12 < v11 {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
							v75 = v71
							v77 = v74
							v78 = v12
							for {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
								if v84 == int32(1) {
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
									v89 = int32(2)
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
									v97 = int32(_a_F_KnownAssignedXidsCompress_0)
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v100 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
									v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v106 = v75 + v100
									v107 = v103
								} else {
									v106 = v75
									v107 = v77
								}
								v109 = v78 + int32(1)
								if v109 != v11 {
									v75 = v106
									v77 = v107
									v78 = v109
									continue
								} else {
									break
								}
								break
							}
							v111 = v106
						} else {
							v111 = v71
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
						v120 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
						if l1 == v120 {
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
							F_LWLockRelease(m, v125+int32(512))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								v134 = m.G0
								v135 = int32(16)
								v136 = v134 - v135
								m.G0 = v136
								F_gettimeofday(m, v136)
								mBase = m.M
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
								v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
								m.G0 = v136 + v135
								*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
								return
							}
						} else {
							v134 = m.G0
							v135 = int32(16)
							v136 = v134 - v135
							m.G0 = v136
							F_gettimeofday(m, v136)
							mBase = m.M
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
							v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
							m.G0 = v136 + v135
							*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
							return
						}
					}
				} else {
					v71 = int32(0)
					if v12 < v11 {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
						v75 = v71
						v77 = v74
						v78 = v12
						for {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
							if v84 == int32(1) {
								v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
								v89 = int32(2)
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
								v97 = int32(_a_F_KnownAssignedXidsCompress_0)
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v100 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
								v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v106 = v75 + v100
								v107 = v103
							} else {
								v106 = v75
								v107 = v77
							}
							v109 = v78 + int32(1)
							if v109 != v11 {
								v75 = v106
								v77 = v107
								v78 = v109
								continue
							} else {
								break
							}
							break
						}
						v111 = v106
					} else {
						v111 = v71
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
					v120 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
					if l1 == v120 {
						v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
						F_LWLockRelease(m, v125+int32(512))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							v134 = m.G0
							v135 = int32(16)
							v136 = v134 - v135
							m.G0 = v136
							F_gettimeofday(m, v136)
							mBase = m.M
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
							v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
							m.G0 = v136 + v135
							*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
							return
						}
					} else {
						v134 = m.G0
						v135 = int32(16)
						v136 = v134 - v135
						m.G0 = v136
						F_gettimeofday(m, v136)
						mBase = m.M
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
						v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
						m.G0 = v136 + v135
						*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
						return
					}
				}
			} else {
				v36 = *(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4]))
				if v36 == int64(0) {
					if l1 == int32(0) {
						v65 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
						v69 = F_LWLockAcquire(m, v65+int32(512), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							v71 = int32(0)
							if v12 < v11 {
								v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v75 = v71
								v77 = v74
								v78 = v12
								for {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
									if v84 == int32(1) {
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
										v89 = int32(2)
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
										v97 = int32(_a_F_KnownAssignedXidsCompress_0)
										v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
										v100 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
										v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
										v106 = v75 + v100
										v107 = v103
									} else {
										v106 = v75
										v107 = v77
									}
									v109 = v78 + int32(1)
									if v109 != v11 {
										v75 = v106
										v77 = v107
										v78 = v109
										continue
									} else {
										break
									}
									break
								}
								v111 = v106
							} else {
								v111 = v71
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
							v120 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
							if l1 == v120 {
								v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
								F_LWLockRelease(m, v125+int32(512))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									v134 = m.G0
									v135 = int32(16)
									v136 = v134 - v135
									m.G0 = v136
									F_gettimeofday(m, v136)
									mBase = m.M
									v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
									v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
									m.G0 = v136 + v135
									*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
									return
								}
							} else {
								v134 = m.G0
								v135 = int32(16)
								v136 = v134 - v135
								m.G0 = v136
								F_gettimeofday(m, v136)
								mBase = m.M
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
								v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
								m.G0 = v136 + v135
								*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
								return
							}
						}
					} else {
						v71 = int32(0)
						if v12 < v11 {
							v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
							v75 = v71
							v77 = v74
							v78 = v12
							for {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
								if v84 == int32(1) {
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
									v89 = int32(2)
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
									v97 = int32(_a_F_KnownAssignedXidsCompress_0)
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v100 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
									v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v106 = v75 + v100
									v107 = v103
								} else {
									v106 = v75
									v107 = v77
								}
								v109 = v78 + int32(1)
								if v109 != v11 {
									v75 = v106
									v77 = v107
									v78 = v109
									continue
								} else {
									break
								}
								break
							}
							v111 = v106
						} else {
							v111 = v71
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
						v120 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
						if l1 == v120 {
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
							F_LWLockRelease(m, v125+int32(512))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								v134 = m.G0
								v135 = int32(16)
								v136 = v134 - v135
								m.G0 = v136
								F_gettimeofday(m, v136)
								mBase = m.M
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
								v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
								m.G0 = v136 + v135
								*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
								return
							}
						} else {
							v134 = m.G0
							v135 = int32(16)
							v136 = v134 - v135
							m.G0 = v136
							F_gettimeofday(m, v136)
							mBase = m.M
							v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
							v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
							m.G0 = v136 + v135
							*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
							return
						}
					}
				} else {
					v42 = m.G0
					v43 = int32(16)
					v44 = v42 - v43
					m.G0 = v44
					F_gettimeofday(m, v44)
					mBase = m.M
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
					v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+8)))
					m.G0 = v44 + v43
					if v48+v47*int64(1000000)-int64(946684800000000) < v36+int64(1000000) {
						return
					} else {
						if l1 == int32(0) {
							v65 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
							v69 = F_LWLockAcquire(m, v65+int32(512), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								v71 = int32(0)
								if v12 < v11 {
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
									v75 = v71
									v77 = v74
									v78 = v12
									for {
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
										if v84 == int32(1) {
											v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
											v89 = int32(2)
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
											v97 = int32(_a_F_KnownAssignedXidsCompress_0)
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
											v100 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
											v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
											v106 = v75 + v100
											v107 = v103
										} else {
											v106 = v75
											v107 = v77
										}
										v109 = v78 + int32(1)
										if v109 != v11 {
											v75 = v106
											v77 = v107
											v78 = v109
											continue
										} else {
											break
										}
										break
									}
									v111 = v106
								} else {
									v111 = v71
								}
								*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
								v120 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
								if l1 == v120 {
									v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
									F_LWLockRelease(m, v125+int32(512))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return
									} else {
										v134 = m.G0
										v135 = int32(16)
										v136 = v134 - v135
										m.G0 = v136
										F_gettimeofday(m, v136)
										mBase = m.M
										v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
										v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
										m.G0 = v136 + v135
										*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
										return
									}
								} else {
									v134 = m.G0
									v135 = int32(16)
									v136 = v134 - v135
									m.G0 = v136
									F_gettimeofday(m, v136)
									mBase = m.M
									v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
									v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
									m.G0 = v136 + v135
									*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
									return
								}
							}
						} else {
							v71 = int32(0)
							if v12 < v11 {
								v74 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
								v75 = v71
								v77 = v74
								v78 = v12
								for {
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v78))))
									if v84 == int32(1) {
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[3]))
										v89 = int32(2)
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v88+v78<<(uint(v89)%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v88+v75<<(uint(v89)%32)))) = v95
										v97 = int32(_a_F_KnownAssignedXidsCompress_0)
										v98 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
										v100 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v98+v75))) = uint8(v100)
										v103 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[2]))
										v106 = v75 + v100
										v107 = v103
									} else {
										v106 = v75
										v107 = v77
									}
									v109 = v78 + int32(1)
									if v109 != v11 {
										v75 = v106
										v77 = v107
										v78 = v109
										continue
									} else {
										break
									}
									break
								}
								v111 = v106
							} else {
								v111 = v71
							}
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v111
							v120 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v120
							if l1 == v120 {
								v125 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[1]))
								F_LWLockRelease(m, v125+int32(512))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									v134 = m.G0
									v135 = int32(16)
									v136 = v134 - v135
									m.G0 = v136
									F_gettimeofday(m, v136)
									mBase = m.M
									v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
									v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
									m.G0 = v136 + v135
									*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
									return
								}
							} else {
								v134 = m.G0
								v135 = int32(16)
								v136 = v134 - v135
								m.G0 = v136
								F_gettimeofday(m, v136)
								mBase = m.M
								v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
								v140 = int64(*(*int32)(unsafe.Add(mBase, uint32(v136)+8)))
								m.G0 = v136 + v135
								*(*int64)(unsafe.Add(mBase, _c_F_KnownAssignedXidsCompress[4])) = v140 + v139*int64(1000000) - int64(946684800000000)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_koi8r_to_iso(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_koi8r_to_iso_0), int32(25), int32(22))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
