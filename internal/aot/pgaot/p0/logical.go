package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogicalDecodingProcessRecord(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v20 != 0 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
		F_ReorderBufferAssignChild(m, v21, v20, v22, v14)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v26 = v25
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+49)))
			v29 = v27 << (uint(int32(5)) % 32)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_LogicalDecodingProcessRecord[0])))
			if v32 == int32(0) {
				F_RmgrNotFound(m, v27)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_LogicalDecodingProcessRecord[1])))
					if v37 != 0 {
						m.T0[v37].(func(*base.Module, int32, int32))(m, l0, v11+int32(8))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v11 + int32(32)
							return
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
						F_ReorderBufferProcessXid(m, v42, v44, v14)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v11 + int32(32)
							return
						}
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_LogicalDecodingProcessRecord[1])))
				if v37 != 0 {
					m.T0[v37].(func(*base.Module, int32, int32))(m, l0, v11+int32(8))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v11 + int32(32)
						return
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
					F_ReorderBufferProcessXid(m, v42, v44, v14)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v11 + int32(32)
						return
					}
				}
			}
		}
	} else {
		v26 = v19
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+49)))
		v29 = v27 << (uint(int32(5)) % 32)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_LogicalDecodingProcessRecord[0])))
		if v32 == int32(0) {
			F_RmgrNotFound(m, v27)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_LogicalDecodingProcessRecord[1])))
				if v37 != 0 {
					m.T0[v37].(func(*base.Module, int32, int32))(m, l0, v11+int32(8))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v11 + int32(32)
						return
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
					F_ReorderBufferProcessXid(m, v42, v44, v14)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v11 + int32(32)
						return
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_LogicalDecodingProcessRecord[1])))
			if v37 != 0 {
				m.T0[v37].(func(*base.Module, int32, int32))(m, l0, v11+int32(8))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					m.G0 = v11 + int32(32)
					return
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
				F_ReorderBufferProcessXid(m, v42, v44, v14)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v11 + int32(32)
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
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
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
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L31
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L27
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v91
	m.G0 = v11 + int32(48)
	return v92
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = v30
	v35 = v28
	goto L13
L11:
	;
	v83 = v28
	goto L12
L12:
	;
	v91 = v83 - l1
	v92 = l1
	goto L9
L13:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_LogicalTapeBackspace[0])))
	if v39 == int64(-1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v83 = v77
	goto L12
L15:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v43 == v44 {
		v91 = int32(0)
		v92 = v35
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = F_BufFileSeekBlock(m, v60, v39)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L22
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_errmsg_internal(m, int32(_a_F_LogicalTapeBackspace_0), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_LogicalTapeBackspace_1), int32(1095), int32(_a_F_LogicalTapeBackspace_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	if v61 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	F_BufFileReadExact(m, v63, v34, int32(_a_F_LogicalTapeBackspace_3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_LogicalTapeBackspace[1])))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v68 != v69 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v39
	v72 = int32(_a_F_LogicalTapeBackspace_4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v72
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_LogicalTapeBackspace[1])))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v74
	v77 = v35 + v72
	if base.Ui32(v77) < base.Ui32(l1) {
		v34 = v67
		v35 = v77
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v39
	F_errmsg(m, int32(_a_F_LogicalTapeBackspace_5), v11+int32(32))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_LogicalTapeBackspace_1), int32(288), int32(_a_F_LogicalTapeBackspace_6))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_LogicalTapeBackspace[1])))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v123
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v39
	F_errmsg_internal(m, int32(_a_F_LogicalTapeBackspace_7), v11)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_LogicalTapeBackspace_1), int32(1106), int32(_a_F_LogicalTapeBackspace_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v65 int64
	_ = v65
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v13 == int32(1) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_LogicalTapeFreeze[0]))) = base.I64_extend_i32_s(int32(0) - v18)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_ltsWriteBlock(m, v22, v23, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v27)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v29 != 0 {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v30 == int32(_a_F_LogicalTapeFreeze_0) {
					v41 = v29
					*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
					if v44 == int64(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v52 = F_BufFileSeekBlock(m, v51, v44)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						if v52 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							F_BufFileReadExact(m, v56, v41, int32(_a_F_LogicalTapeFreeze_0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
								if v62 < int64(0) {
									v65 = int64(-1)
								} else {
									v65 = v62
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v65
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
								if int64(0) <= v69 {
									v74 = int32(_a_F_LogicalTapeFreeze_1)
								} else {
									v74 = int32(0) - base.I32_wrap_i64(v69)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v74
								if l1 != 0 {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
									if v77 == int32(1) {
										F_BufFileDumpBuffer(m, v76)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											v82 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
											v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v82 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
										v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
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
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
									F_errmsg(m, int32(_a_F_LogicalTapeFreeze_2), v10)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_LogicalTapeFreeze_3), int32(288), int32(_a_F_LogicalTapeFreeze_4))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
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
					F_pfree(m, v29)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v36 = F_palloc(m, int32(_a_F_LogicalTapeFreeze_0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(_a_F_LogicalTapeFreeze_0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36
							v41 = v36
							*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
							v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
							if v44 == int64(-1) {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
							} else {
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							v52 = F_BufFileSeekBlock(m, v51, v44)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								if v52 == int32(0) {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									F_BufFileReadExact(m, v56, v41, int32(_a_F_LogicalTapeFreeze_0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
										if v62 < int64(0) {
											v65 = int64(-1)
										} else {
											v65 = v62
										}
										*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v65
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
										if int64(0) <= v69 {
											v74 = int32(_a_F_LogicalTapeFreeze_1)
										} else {
											v74 = int32(0) - base.I32_wrap_i64(v69)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v74
										if l1 != 0 {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
											if v77 == int32(1) {
												F_BufFileDumpBuffer(m, v76)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													v82 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
													v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
													m.G0 = v10 + int32(16)
													return
												}
											} else {
												v82 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
												v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
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
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
											F_errmsg(m, int32(_a_F_LogicalTapeFreeze_2), v10)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_LogicalTapeFreeze_3), int32(288), int32(_a_F_LogicalTapeFreeze_4))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
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
				v36 = F_palloc(m, int32(_a_F_LogicalTapeFreeze_0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(_a_F_LogicalTapeFreeze_0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36
					v41 = v36
					*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
					if v44 == int64(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					v52 = F_BufFileSeekBlock(m, v51, v44)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						if v52 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							F_BufFileReadExact(m, v56, v41, int32(_a_F_LogicalTapeFreeze_0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
								if v62 < int64(0) {
									v65 = int64(-1)
								} else {
									v65 = v62
								}
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v65
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
								if int64(0) <= v69 {
									v74 = int32(_a_F_LogicalTapeFreeze_1)
								} else {
									v74 = int32(0) - base.I32_wrap_i64(v69)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v74
								if l1 != 0 {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
									if v77 == int32(1) {
										F_BufFileDumpBuffer(m, v76)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											v82 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
											v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										v82 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
										v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
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
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
									F_errmsg(m, int32(_a_F_LogicalTapeFreeze_2), v10)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_LogicalTapeFreeze_3), int32(288), int32(_a_F_LogicalTapeFreeze_4))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
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
		v27 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v27)
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v29 != 0 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v30 == int32(_a_F_LogicalTapeFreeze_0) {
				v41 = v29
				*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
				if v44 == int64(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
				} else {
				}
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v52 = F_BufFileSeekBlock(m, v51, v44)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					if v52 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						F_BufFileReadExact(m, v56, v41, int32(_a_F_LogicalTapeFreeze_0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
							if v62 < int64(0) {
								v65 = int64(-1)
							} else {
								v65 = v62
							}
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v65
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
							if int64(0) <= v69 {
								v74 = int32(_a_F_LogicalTapeFreeze_1)
							} else {
								v74 = int32(0) - base.I32_wrap_i64(v69)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v74
							if l1 != 0 {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
								if v77 == int32(1) {
									F_BufFileDumpBuffer(m, v76)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
										v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v82 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
									v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
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
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
								F_errmsg(m, int32(_a_F_LogicalTapeFreeze_2), v10)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_LogicalTapeFreeze_3), int32(288), int32(_a_F_LogicalTapeFreeze_4))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
				F_pfree(m, v29)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = F_palloc(m, int32(_a_F_LogicalTapeFreeze_0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(_a_F_LogicalTapeFreeze_0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36
						v41 = v36
						*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
						v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
						if v44 == int64(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
						} else {
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						v52 = F_BufFileSeekBlock(m, v51, v44)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							if v52 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
								F_BufFileReadExact(m, v56, v41, int32(_a_F_LogicalTapeFreeze_0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
									if v62 < int64(0) {
										v65 = int64(-1)
									} else {
										v65 = v62
									}
									*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v65
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
									if int64(0) <= v69 {
										v74 = int32(_a_F_LogicalTapeFreeze_1)
									} else {
										v74 = int32(0) - base.I32_wrap_i64(v69)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v74
									if l1 != 0 {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
										if v77 == int32(1) {
											F_BufFileDumpBuffer(m, v76)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												v82 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
												v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
												m.G0 = v10 + int32(16)
												return
											}
										} else {
											v82 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
											v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
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
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
										F_errmsg(m, int32(_a_F_LogicalTapeFreeze_2), v10)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_LogicalTapeFreeze_3), int32(288), int32(_a_F_LogicalTapeFreeze_4))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
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
			v36 = F_palloc(m, int32(_a_F_LogicalTapeFreeze_0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(_a_F_LogicalTapeFreeze_0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v36
				v41 = v36
				*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
				v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v44
				if v44 == int64(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(-1)
				} else {
				}
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v52 = F_BufFileSeekBlock(m, v51, v44)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					if v52 == int32(0) {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						F_BufFileReadExact(m, v56, v41, int32(_a_F_LogicalTapeFreeze_0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
							if v62 < int64(0) {
								v65 = int64(-1)
							} else {
								v65 = v62
							}
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v65
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_LogicalTapeFreeze[0])))
							if int64(0) <= v69 {
								v74 = int32(_a_F_LogicalTapeFreeze_1)
							} else {
								v74 = int32(0) - base.I32_wrap_i64(v69)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v74
							if l1 != 0 {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
								if v77 == int32(1) {
									F_BufFileDumpBuffer(m, v76)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
										v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v82 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)) = uint8(v82)
									v84 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = v84
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
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v10))) = v44
								F_errmsg(m, int32(_a_F_LogicalTapeFreeze_2), v10)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_LogicalTapeFreeze_3), int32(288), int32(_a_F_LogicalTapeFreeze_4))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
