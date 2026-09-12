package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogicalDecodingProcessRecord(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v19 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
		F_ReorderBufferAssignChild(m, v20, v19, v21, v13)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v25 = v24
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+49)))
			v28 = v26 << (uint(int32(5)) % 32)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[510])))
			if v31 == int32(0) {
				F_RmgrNotFound(m, v26)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[511])))
					if v36 != 0 {
						m.T0[v36].(func(*base.Module, int32, int32))(m, l0, v10+int32(8))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
						F_ReorderBufferProcessXid(m, v41, v43, v13)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[511])))
				if v36 != 0 {
					m.T0[v36].(func(*base.Module, int32, int32))(m, l0, v10+int32(8))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v10 + int32(32)
						return
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
					F_ReorderBufferProcessXid(m, v41, v43, v13)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v10 + int32(32)
						return
					}
				}
			}
		}
	} else {
		v25 = v18
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+49)))
		v28 = v26 << (uint(int32(5)) % 32)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[510])))
		if v31 == int32(0) {
			F_RmgrNotFound(m, v26)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[511])))
				if v36 != 0 {
					m.T0[v36].(func(*base.Module, int32, int32))(m, l0, v10+int32(8))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v10 + int32(32)
						return
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
					F_ReorderBufferProcessXid(m, v41, v43, v13)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						m.G0 = v10 + int32(32)
						return
					}
				}
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[511])))
			if v36 != 0 {
				m.T0[v36].(func(*base.Module, int32, int32))(m, l0, v10+int32(8))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					m.G0 = v10 + int32(32)
					return
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+36))
				F_ReorderBufferProcessXid(m, v41, v43, v13)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					m.G0 = v10 + int32(32)
					return
				}
			}
		}
	}
}
func F_LogicalTapeBackspace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = F_palloc(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.Ui32(v28) < base.Ui32(l1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v24
	v26 = F_ltsReadFillBuffer(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L31
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L27
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v95
	m.G0 = v11 + int32(48)
	return v96
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = v30
	v35 = v28
	goto L13
L11:
	;
	v87 = v28
	goto L12
L12:
	;
	v95 = v87 - l1
	v96 = l1
	goto L9
L13:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_consts[986])))
	if v41 == int64(-1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v87 = v81
	goto L12
L15:
	;
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v45 == v46 {
		v95 = int32(0)
		v96 = v35
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = F_BufFileSeekBlock(m, v62, v41)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L22
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errmsg_internal(m, int32(371892), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(498793), int32(1095), int32(418342))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	if v63 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	F_BufFileReadExact(m, v65, v34, int32(8192))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[985])))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v72 != v73 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v41
	v76 = int32(8176)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v69)+uint32(_consts[985])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v78
	v81 = v35 + v76
	if base.Ui32(v81) < base.Ui32(l1) {
		v34 = v69
		v35 = v81
		goto L13
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v41
	F_errmsg(m, int32(386992), v11+int32(32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(498793), int32(288), int32(317160))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v126)+uint32(_consts[985])))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v41
	F_errmsg_internal(m, int32(430191), v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(498793), int32(1106), int32(418342))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LogicalTapeFreeze(m *base.Module, l0 int32, l1 int32) {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v13 == int32(1) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v16)+uint32(_consts[985]))) = base.I64_extend_i32_s(int32(0) - v20)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_ltsWriteBlock(m, v24, v25, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v29)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v31 != 0 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v32 == int32(8192) {
					v43 = v31
					*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v46
					if v46 == int64(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
					} else {
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					v54 = F_BufFileSeekBlock(m, v53, v46)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						if v54 == int32(0) {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							F_BufFileReadExact(m, v58, v43, int32(8192))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
								if v66 < int64(0) {
									v69 = int64(-1)
								} else {
									v69 = v66
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v69
								v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
								if int64(0) <= v73 {
									v78 = int32(8176)
								} else {
									v78 = int32(0) - base.I32_wrap_i64(v73)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v78
								if l1 != 0 {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
									if v81 == int32(1) {
										F_BufFileDumpBuffer(m, v80)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											v86 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
											v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v86 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
										v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v46
									F_errmsg(m, int32(386992), v10)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										F_errfinish(m, int32(498793), int32(288), int32(317160))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_pfree(m, v31)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v38 = F_palloc(m, int32(8192))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(8192)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
							v43 = v38
							*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
							v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v46
							if v46 == int64(-1) {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
							} else {
							}
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							v54 = F_BufFileSeekBlock(m, v53, v46)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								if v54 == int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									F_BufFileReadExact(m, v58, v43, int32(8192))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
										if v66 < int64(0) {
											v69 = int64(-1)
										} else {
											v69 = v66
										}
										*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v69
										v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
										if int64(0) <= v73 {
											v78 = int32(8176)
										} else {
											v78 = int32(0) - base.I32_wrap_i64(v73)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v78
										if l1 != 0 {
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
											if v81 == int32(1) {
												F_BufFileDumpBuffer(m, v80)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v86 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
													v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
													m.G0 = v10 + int32(16)
													return
												}
											} else {
												v86 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
												v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
												m.G0 = v10 + int32(16)
												return
											}
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v46
											F_errmsg(m, int32(386992), v10)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												F_errfinish(m, int32(498793), int32(288), int32(317160))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
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
			} else {
				v38 = F_palloc(m, int32(8192))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(8192)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
					v43 = v38
					*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v46
					if v46 == int64(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
					} else {
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					v54 = F_BufFileSeekBlock(m, v53, v46)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						if v54 == int32(0) {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							F_BufFileReadExact(m, v58, v43, int32(8192))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
								if v66 < int64(0) {
									v69 = int64(-1)
								} else {
									v69 = v66
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v69
								v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
								if int64(0) <= v73 {
									v78 = int32(8176)
								} else {
									v78 = int32(0) - base.I32_wrap_i64(v73)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v78
								if l1 != 0 {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
									if v81 == int32(1) {
										F_BufFileDumpBuffer(m, v80)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											v86 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
											v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v86 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
										v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v46
									F_errmsg(m, int32(386992), v10)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										F_errfinish(m, int32(498793), int32(288), int32(317160))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
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
	} else {
		v29 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v29)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v31 != 0 {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v32 == int32(8192) {
				v43 = v31
				*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v46
				if v46 == int64(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
				} else {
				}
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				v54 = F_BufFileSeekBlock(m, v53, v46)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					if v54 == int32(0) {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						F_BufFileReadExact(m, v58, v43, int32(8192))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
							if v66 < int64(0) {
								v69 = int64(-1)
							} else {
								v69 = v66
							}
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v69
							v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
							if int64(0) <= v73 {
								v78 = int32(8176)
							} else {
								v78 = int32(0) - base.I32_wrap_i64(v73)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v78
							if l1 != 0 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
								if v81 == int32(1) {
									F_BufFileDumpBuffer(m, v80)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										v86 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
										v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
									v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v10))) = v46
								F_errmsg(m, int32(386992), v10)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									F_errfinish(m, int32(498793), int32(288), int32(317160))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_pfree(m, v31)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v38 = F_palloc(m, int32(8192))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(8192)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
						v43 = v38
						*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
						v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v46
						if v46 == int64(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
						} else {
						}
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						v54 = F_BufFileSeekBlock(m, v53, v46)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							if v54 == int32(0) {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								F_BufFileReadExact(m, v58, v43, int32(8192))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
									if v66 < int64(0) {
										v69 = int64(-1)
									} else {
										v69 = v66
									}
									*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v69
									v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
									if int64(0) <= v73 {
										v78 = int32(8176)
									} else {
										v78 = int32(0) - base.I32_wrap_i64(v73)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v78
									if l1 != 0 {
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
										if v81 == int32(1) {
											F_BufFileDumpBuffer(m, v80)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v86 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
												v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
												m.G0 = v10 + int32(16)
												return
											}
										} else {
											v86 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
											v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v46
										F_errmsg(m, int32(386992), v10)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											F_errfinish(m, int32(498793), int32(288), int32(317160))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
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
		} else {
			v38 = F_palloc(m, int32(8192))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(8192)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
				v43 = v38
				*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v46
				if v46 == int64(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
				} else {
				}
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				v54 = F_BufFileSeekBlock(m, v53, v46)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					if v54 == int32(0) {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						F_BufFileReadExact(m, v58, v43, int32(8192))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
							if v66 < int64(0) {
								v69 = int64(-1)
							} else {
								v69 = v66
							}
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v69
							v73 = *(*int64)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[985])))
							if int64(0) <= v73 {
								v78 = int32(8176)
							} else {
								v78 = int32(0) - base.I32_wrap_i64(v73)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v78
							if l1 != 0 {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
								if v81 == int32(1) {
									F_BufFileDumpBuffer(m, v80)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										v86 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
										v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v86 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)) = uint8(v86)
									v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v10))) = v46
								F_errmsg(m, int32(386992), v10)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									F_errfinish(m, int32(498793), int32(288), int32(317160))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
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
