package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_Aio(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+20)))
	if v4 != int32(1) {
		return
	} else {
		v7 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+20)) = uint8(v7)
		F_pgaio_submit_staged(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v13 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				if v13 == int32(0) {
					return
				} else {
					F_errmsg_internal(m, int32(252263), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						F_errfinish(m, int32(487199), int32(1206), int32(237022))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_AtEOXact_Enum(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[117])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[116])) = v2
	return
}
func F_AtEOXact_Namespace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	if l1 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, _consts[171]))
		if v4 == int32(0) {
			return
		} else {
			if l0 != 0 {
				F_before_shmem_exit(m, int32(469), int32(0))
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[171])) = int32(0)
					return
				}
			} else {
				v12 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _consts[423])) = uint8(v12)
				v15 = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[127])) = v15
				*(*int32)(unsafe.Add(mBase, _consts[126])) = v15
				*(*uint8)(unsafe.Add(mBase, _consts[424])) = uint8(v15)
				v24 = *(*int32)(unsafe.Add(mBase, _consts[128]))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v15
				*(*int32)(unsafe.Add(mBase, _consts[171])) = int32(0)
				return
			}
		}
	}
}
func F_AtEOXact_Parallel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v4 == int32(4079932) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = v4
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	F_DestroyParallelContext(m, v10)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L13
	}
L7:
	;
	v15 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v15 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errmsg_internal(m, int32(60192), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(488530), int32(1290), int32(302609))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	if v31 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v31 != int32(4079932) {
		v10 = v31
		goto L4
	} else {
		goto L15
	}
L15:
	;
	goto L5
}
func F_AtEOXact_PgStat(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v196 int32
	_ = v196
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[927]))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v15 = int32(4450160)
	goto L6
L5:
	;
	v15 = int32(4450164)
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16 + int32(1)
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[927])) = int32(0)
	F_pgstat_clear_snapshot(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L38
	} else {
		goto L44
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v25
	goto L12
L10:
	;
	goto L11
L11:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v116 == int32(0) {
		goto L7
	} else {
		goto L26
	}
L12:
	;
	if l0 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+40))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+40)) = v46 + v47
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+48)) = v50 + v51
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v45)+56))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+56)) = v54 + v55
	if l0 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	if v36 != int32(1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v27)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v27)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v27)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v43
	goto L14
L17:
	;
	v99 = v98 + v45
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = v96 + v100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v105 != 0 {
		v27 = v105
		goto L12
	} else {
		goto L25
	}
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+80)) = uint8(v58)
	if v58 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	v96 = v91 + v92
	v98 = int32(96)
	goto L17
L21:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+88)) = v75 - v76 + v73
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+96)) = v80 + v81 + v74
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	v96 = v85 + (v86 + v87)
	v98 = int32(104)
	goto L17
L22:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v45)+88))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v45)+96))
	v73 = v62
	v74 = v63
	goto L21
L23:
	;
	goto L24
L24:
	;
	v65 = v45 + int32(88)
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
	v73 = v66
	v74 = v66
	goto L21
L25:
	;
	goto L13
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v119 == int32(0) {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v123 = v22 + int32(8)
	if v119 == v123 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v127 = v119
	v131 = v3
	goto L29
L29:
	;
	v136 = v127 - int32(20)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127-int32(4)))))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v127-int32(12))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if l0 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	if v160 <= int32(0) {
		goto L7
	} else {
		goto L42
	}
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v161
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v166 - int32(1)
	F_pfree(m, v136)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L38
	} else {
		goto L40
	}
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v127-int32(16))))
	v154 = F_pgstat_drop_entry(m, v150, v153, v142)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	if v139&int32(1) == int32(0) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v139&int32(1) != 0 {
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v160 = v131
	v161 = v143
	goto L31
L37:
	;
	v160 = v131
	v161 = v143
	goto L31
L38:
	;
	return
L39:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v160 = v131 + (v154 ^ int32(1))
	v161 = v159
	goto L31
L40:
	;
	if v143 != v123 {
		v127 = v143
		v131 = v160
		goto L29
	} else {
		goto L41
	}
L41:
	;
	goto L30
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[267]))
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v177)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v177)+16)) = v178 + int64(1)
	goto L43
L43:
	;
	goto L7
L44:
	;
	return
}
func F_AtEOXact_on_commit_actions(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	if v6 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = v2
	v11 = v6
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v13 <= v10 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15+v10<<(uint(int32(2))%32))))
	if l0 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(0)
	if v11 != 0 {
		v10 = v10 + int32(1)
		goto L3
	} else {
		goto L17
	}
L7:
	;
	v24 = int32(4367436)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v27 = F_list_delete_nth_cell(m, v26, v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v20 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v21 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	goto L6
L12:
	;
	goto L7
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[495])) = v27
	F_pfree(m, v19)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v27 != 0 {
		v11 = v27
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	goto L4
}
