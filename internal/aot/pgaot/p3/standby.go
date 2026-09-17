package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StandbyReleaseLockTree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(0) < l1 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_StandbyReleaseLockTree[0]))
	v15 = int32(0)
	v17 = F_hash_search(m, v12, v8+int32(8), v15, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_StandbyReleaseAllLocks(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L10
	}
L5:
	;
	return
L6:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_StandbyReleaseXidEntryLocks(m, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_StandbyReleaseLockTree[0]))
	v27 = F_hash_search(m, v24, v17, int32(2), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	goto L1
L11:
	;
	v35 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	m.G0 = v8 + int32(16)
	return
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2+v35<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v43
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	v67 = v35 + int32(1)
	if v67 != l1 {
		v35 = v67
		goto L14
	} else {
		goto L25
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_StandbyReleaseLockTree[0]))
	v49 = int32(0)
	v51 = F_hash_search(m, v46, v8+int32(12), v49, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_StandbyReleaseAllLocks(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L24
	}
L20:
	;
	if v51 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_StandbyReleaseXidEntryLocks(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_StandbyReleaseLockTree[0]))
	v61 = F_hash_search(m, v58, v51, int32(2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	goto L16
L25:
	;
	goto L15
}
func F_StandbyReleaseXidEntryLocks(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v10
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	m.G0 = v8 + int32(48)
	return
L4:
	;
	v18 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v20
	F_errmsg_internal(m, int32(_a_F_StandbyReleaseXidEntryLocks_0), v8+int32(16))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v37
	v45 = F_LockRelease(m, v8+int32(32), int32(8), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	F_errfinish(m, int32(_a_F_StandbyReleaseXidEntryLocks_1), int32(1046), int32(_a_F_StandbyReleaseXidEntryLocks_2))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_StandbyReleaseXidEntryLocks[0]))
	v71 = F_hash_search(m, v68, v14, int32(2), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L20
	}
L14:
	;
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v49 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	if v49 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v53
	F_errmsg_internal(m, int32(_a_F_StandbyReleaseXidEntryLocks_3), v8)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_StandbyReleaseXidEntryLocks_1), int32(1053), int32(_a_F_StandbyReleaseXidEntryLocks_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	if v66 != 0 {
		v14 = v66
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L5
}
func F_StandbySlotsHaveCaughtup(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v257 int64
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[0]))
	if v18 == v3 {
		v275 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v275
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[1])))
	if v23 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v33 != 0 {
		v275 = v16
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[2]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+316))
	v31 = base.B2i32(v29 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[1])) = uint8(v31)
	v33 = v31
	goto L6
L5:
	;
	v33 = int32(0)
	goto L6
L6:
	;
	goto L3
L7:
	;
	v35 = *(*int64)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[3]))
	if base.B2i32(v35 != int64(0))&base.B2i32(base.Ui64(l0) <= base.Ui64(v35)) != 0 {
		v275 = v16
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[4]))
	v45 = F_LWLockAcquire(m, v41+int32(_a_F_StandbySlotsHaveCaughtup_0), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[4]))
	F_LWLockRelease(m, v261+int32(_a_F_StandbySlotsHaveCaughtup_0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L67
	}
L12:
	;
	v255 = v3
	v257 = int64(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v63 = v50 + int32(4)
	v65 = v3
	v67 = int64(0)
	goto L15
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[5]))
	if int32(0) < v70 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v255 = v243
	v257 = v237
	goto L11
L17:
	;
	if base.Ui64(v67-int64(1)) < base.Ui64(v179) {
		goto L63
	} else {
		goto L64
	}
L18:
	;
	F_errcode(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L58
	}
L19:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	if v147 != 0 {
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[6]))
	v78 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	v138 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L37
	}
L23:
	;
	v89 = v75 + v78*int32(288)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v90 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	v94 = v89 + int32(24)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if base.B2i32(v97 == int32(0))|base.B2i32(v97 != v100) != 0 {
		v118 = v97
		v119 = v100
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v124 = v78 + int32(1)
	if v124 != v70 {
		v78 = v124
		goto L23
	} else {
		goto L36
	}
L28:
	;
	if v118-v119 == int32(0) {
		goto L19
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v103 = v63
	v104 = v94
	goto L31
L31:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v108 == int32(0) {
		v118 = v108
		v119 = v107
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v118 = v108
	v119 = v107
	goto L29
L33:
	;
	v111 = int32(1)
	if v108 == v107 {
		v103 = v103 + v111
		v104 = v104 + v111
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L27
L36:
	;
	goto L24
L37:
	;
	if v138 == int32(0) {
		v255 = v65
		v257 = v67
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v200 = int32(_a_F_StandbySlotsHaveCaughtup_1)
	v201 = int32(2958)
	v202 = int32(_a_F_StandbySlotsHaveCaughtup_2)
	v205 = int32(_a_F_StandbySlotsHaveCaughtup_3)
	v210 = int32(50856066)
	goto L18
L39:
	;
	v149 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(1)
	if v158 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v149 == int32(0) {
		v255 = v65
		v257 = v67
		goto L11
	} else {
		goto L43
	}
L43:
	;
	v200 = int32(_a_F_StandbySlotsHaveCaughtup_4)
	v201 = int32(2972)
	v202 = int32(_a_F_StandbySlotsHaveCaughtup_5)
	v205 = int32(_a_F_StandbySlotsHaveCaughtup_6)
	v210 = int32(50856066)
	goto L18
L44:
	;
	F_s_lock(m, v89, int32(_a_F_StandbySlotsHaveCaughtup_7), int32(2976), int32(_a_F_StandbySlotsHaveCaughtup_8))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v89)+112))
	if v168 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L46
L48:
	;
	v200 = v193
	v201 = v197
	v202 = v194
	v205 = v195
	v210 = int32(325)
	goto L18
L49:
	;
	v170 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v89)+104))
	if base.B2i32(v179 != int64(0))&base.B2i32(base.Ui64(l0) <= base.Ui64(v179)) != 0 {
		goto L17
	} else {
		goto L54
	}
L52:
	;
	if v170 == int32(0) {
		v255 = v65
		v257 = v67
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v193 = int32(_a_F_StandbySlotsHaveCaughtup_9)
	v194 = int32(_a_F_StandbySlotsHaveCaughtup_2)
	v195 = int32(_a_F_StandbySlotsHaveCaughtup_10)
	v197 = int32(2992)
	goto L48
L54:
	;
	if v178 != 0 {
		v255 = v65
		v257 = v67
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v185 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	if v185 == int32(0) {
		v255 = v65
		v257 = v67
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v193 = int32(_a_F_StandbySlotsHaveCaughtup_11)
	v194 = int32(_a_F_StandbySlotsHaveCaughtup_2)
	v195 = int32(_a_F_StandbySlotsHaveCaughtup_12)
	v197 = int32(3007)
	goto L48
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = int32(_a_F_StandbySlotsHaveCaughtup_13)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v63
	F_errmsg(m, v205, v14+int32(32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v63
	F_errdetail(m, v202, v14+int32(16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(_a_F_StandbySlotsHaveCaughtup_13)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v63
	F_errhint(m, v200, v14)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_StandbySlotsHaveCaughtup_7), v201, int32(_a_F_StandbySlotsHaveCaughtup_8))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v255 = v65
	v257 = v67
	goto L11
L63:
	;
	v237 = v67
	goto L65
L64:
	;
	v237 = v179
	goto L65
L65:
	;
	v238 = F_strlen(m, v63)
	mBase = m.M
	v240 = int32(1)
	v243 = v65 + v240
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[0]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	if v243 < v246 {
		v63 = v238 + v63 + v240
		v65 = v243
		v67 = v237
		goto L15
	} else {
		goto L66
	}
L66:
	;
	goto L16
L67:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[0]))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v255 != v268 {
		v275 = int32(0)
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_StandbySlotsHaveCaughtup[3])) = v257
	v275 = int32(1)
	goto L1
}
