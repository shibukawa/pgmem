package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_apply_handle_commit_internal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v6 = *(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[0]))
	if v6 == int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[1]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	goto L17
L2:
	;
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[0]))
	if v14 != int64(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v11 + int32(16)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[1]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	goto L13
L6:
	;
	return
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = *(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[0]))
	*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v22)
	v25 = int64(base.Ui64(v22) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v11))) = uint32(v25)
	F_errmsg(m, int32(_a_F_apply_handle_commit_internal_1), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[0])) = int64(0)
	goto L5
L11:
	;
	F_errfinish(m, int32(_a_F_apply_handle_commit_internal_2), int32(_a_F_apply_handle_commit_internal_3), int32(_a_F_apply_handle_commit_internal_4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v45 == int32(2) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[6])) = uint8(v129)
	return
L17:
	;
	if v54 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_clear_subscription_skip_lsn(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L39
	}
L21:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[2])) = v61
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[3])) = v64
	F_CommitTransactionCommand(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[1]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	goto L23
L23:
	;
	if base.Ui32(int32(1)) < base.Ui32(v70) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = F_EndTransactionBlock(m, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v79 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[4]))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[5]))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+16)))
	if v86 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v89 == int32(4) {
		goto L16
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[8])) = v94
	v97 = F_palloc(m, int32(24))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+16)) = v83
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = v82
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[9]))
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v109
	v111 = int32(_a_F_apply_handle_commit_internal_0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v97
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[10])) = v97
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[8])) = v118
	goto L16
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[10]))
	v109 = v104
	goto L35
L37:
	;
	goto L38
L38:
	;
	v106 = int32(_a_F_apply_handle_commit_internal_0)
	*(*int32)(unsafe.Add(mBase, _c_F_apply_handle_commit_internal[9])) = v106
	v109 = v106
	goto L35
L39:
	;
	F_maybe_reread_subscription(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	goto L16
}
func F_apply_handle_delete_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	F_EvalPlanQualInit(m, v10+int32(44), v13, v5, v5, int32(-1), v5)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_TargetPrivilegesCheck(m, v12, int64(2))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			v34 = F_table_slot_create(m, v12, v28+int32(104))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if l3 != 0 {
					v36 = F_RelationFindReplTupleByIndex(m, v12, l3, l2, v34)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 != 0 {
							v48 = F_GetTupleTransactionInfo(m, v34, v10+int32(24), v10+int32(28), v10+int32(32))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								if v48 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									F_TargetPrivilegesCheck(m, v71, int64(8))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_EvalPlanQualEnd(m, v10+int32(44))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v10 + int32(96)
												return
											}
										}
									}
								} else {
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)))
									v54 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_handle_delete_internal[0])))
									if v52 == v54 {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										F_TargetPrivilegesCheck(m, v71, int64(8))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_EvalPlanQualEnd(m, v10+int32(44))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													m.G0 = v10 + int32(96)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v34
										v58 = v10 + int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v58
										v65 = F_list_make1_impl(m, int32(1), v10)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_ReportApplyConflict(m, v13, l1, int32(15), int32(5), l2, int32(0), v65)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												F_TargetPrivilegesCheck(m, v71, int64(8))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														F_EvalPlanQualEnd(m, v10+int32(44))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															m.G0 = v10 + int32(96)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v80 = v10 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v80
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v80
							v89 = F_list_make1_impl(m, int32(1), v10+int32(4))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								F_ReportApplyConflict(m, v13, l1, int32(15), int32(6), l2, int32(0), v89)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_EvalPlanQualEnd(m, v10+int32(44))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										m.G0 = v10 + int32(96)
										return
									}
								}
							}
						}
					}
				} else {
					v38 = F_RelationFindReplTupleSeq(m, v12, l2, v34)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						if v38 == int32(0) {
							v80 = v10 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v80
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v80
							v89 = F_list_make1_impl(m, int32(1), v10+int32(4))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								F_ReportApplyConflict(m, v13, l1, int32(15), int32(6), l2, int32(0), v89)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_EvalPlanQualEnd(m, v10+int32(44))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										m.G0 = v10 + int32(96)
										return
									}
								}
							}
						} else {
							v48 = F_GetTupleTransactionInfo(m, v34, v10+int32(24), v10+int32(28), v10+int32(32))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								if v48 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									F_TargetPrivilegesCheck(m, v71, int64(8))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_EvalPlanQualEnd(m, v10+int32(44))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												m.G0 = v10 + int32(96)
												return
											}
										}
									}
								} else {
									v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)))
									v54 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_apply_handle_delete_internal[0])))
									if v52 == v54 {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										F_TargetPrivilegesCheck(m, v71, int64(8))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_EvalPlanQualEnd(m, v10+int32(44))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													m.G0 = v10 + int32(96)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v34
										v58 = v10 + int32(16)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v58
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v58
										v65 = F_list_make1_impl(m, int32(1), v10)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_ReportApplyConflict(m, v13, l1, int32(15), int32(5), l2, int32(0), v65)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v34
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												F_TargetPrivilegesCheck(m, v71, int64(8))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return
												} else {
													F_ExecSimpleRelationDelete(m, l1, v13, v10+int32(44), v34)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														F_EvalPlanQualEnd(m, v10+int32(44))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															m.G0 = v10 + int32(96)
															return
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
			}
		}
	}
}
func F_apply_handle_prepare_internal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = m.G0
	v5 = v3 - int32(208)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_prepare_internal[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_TwoPhaseTransactionGid(m, v9, v10, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_apply_handle_prepare_internal[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
		if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v15)) == int32(0) {
			F_BeginTransactionBlock(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_CommitTransactionCommand(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_prepare_internal[2])) = v25
					v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_prepare_internal[3])) = v28
					v30 = F_PrepareTransactionBlock(m, v5)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v5 + int32(208)
						return
					}
				}
			}
		} else {
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_prepare_internal[2])) = v25
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int64)(unsafe.Add(mBase, _c_F_apply_handle_prepare_internal[3])) = v28
			v30 = F_PrepareTransactionBlock(m, v5)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v5 + int32(208)
				return
			}
		}
	}
}
func F_attach_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = F_palloc(m, int32(656))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_attach_internal[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v11
		v15 = int32(0)
		base.MemoryFill(m, v5+int32(28), v15, int32(624))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = l0 + int32(2048)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = l0 + int32(1496)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
		v30 = F_LWLockAcquire(m, l0+int32(1476), v15)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1460))
			if v32 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_attach_internal_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_attach_internal_1), int32(1414), int32(_a_F_attach_internal_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+1460)) = v32 + int32(1)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+1468))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+652)) = v55
				F_LWLockRelease(m, v54+int32(1476))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					return v5
				}
			}
		}
	}
}
func F_internal_bpchar_pattern_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	v7 = int32(1)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v11 = v9 & v7
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v7
	goto L3
L2:
	;
	v12 = int32(4)
	goto L3
L3:
	;
	v13 = v12 + l0
	if v9 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = v40
	goto L15
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v19 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v30 = int32(1)
	if v11 != 0 {
		v40 = int32(base.Ui32(v9)>>(uint(v30)%32)) - v30
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v22 = int32(16)
	goto L10
L9:
	;
	v22 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = int32(4)
	goto L13
L12:
	;
	v29 = v22
	goto L13
L13:
	;
	v40 = v29
	goto L4
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v46 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v60 = int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v64 = v62 & v60
	if v64 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v58 = v40 & (v40 >> (uint(int32(31)) % 32))
	goto L17
L19:
	;
	goto L20
L20:
	;
	v53 = v46 - int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v53))))
	if v55 == int32(32) {
		v46 = v53
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v58 = v46
	goto L17
L22:
	;
	v65 = v60
	goto L24
L23:
	;
	v65 = int32(4)
	goto L24
L24:
	;
	v66 = v65 + l1
	if v62 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v99 = v93
	goto L36
L26:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v72 == int32(18) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v83 = int32(1)
	if v64 != 0 {
		v93 = int32(base.Ui32(v62)>>(uint(v83)%32)) - v83
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v75 = int32(16)
	goto L31
L30:
	;
	v75 = int32(0)
	goto L31
L31:
	;
	if base.Ui32((v72-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v82 = int32(4)
	goto L34
L33:
	;
	v82 = v75
	goto L34
L34:
	;
	v93 = v82
	goto L25
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v93 = int32(base.Ui32(v87)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	if v99 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v113 = base.B2i32(v58 < v111)
	if v58 < v111 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	v111 = v93 & (v93 >> (uint(int32(31)) % 32))
	goto L38
L40:
	;
	goto L41
L41:
	;
	v106 = v99 - int32(1)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v106))))
	if v108 == int32(32) {
		v99 = v106
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v111 = v99
	goto L38
L43:
	;
	return v179
L44:
	;
	v114 = v58
	goto L46
L45:
	;
	v114 = v111
	goto L46
L46:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v114) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	if v176 != 0 {
		v179 = v176
		goto L43
	} else {
		goto L65
	}
L48:
	;
	v176 = int32(0)
	goto L47
L49:
	;
	v150 = v145
	v151 = v146
	v152 = v147
	goto L59
L50:
	;
	if (v13|v66)&int32(3) != 0 {
		v145 = v13
		v146 = v66
		v147 = v114
		goto L49
	} else {
		goto L53
	}
L51:
	;
	v138 = v13
	v139 = v66
	v140 = v114
	goto L52
L52:
	;
	if v140 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L53:
	;
	v122 = v13
	v123 = v66
	v124 = v114
	goto L54
L54:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v127 != v128 {
		v145 = v122
		v146 = v123
		v147 = v124
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v138 = v133
	v139 = v131
	v140 = v135
	goto L52
L56:
	;
	v130 = int32(4)
	v131 = v123 + v130
	v133 = v122 + v130
	v135 = v124 - v130
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v122 = v133
		v123 = v131
		v124 = v135
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v145 = v138
	v146 = v139
	v147 = v140
	goto L49
L59:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 == v156 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v176 = v155 - v156
	goto L47
L61:
	;
	v158 = int32(1)
	v163 = v152 - v158
	if v163 != 0 {
		v150 = v150 + v158
		v151 = v151 + v158
		v152 = v163
		goto L59
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L48
L65:
	;
	if v58 < v111 {
		v179 = int32(-1)
		goto L43
	} else {
		goto L66
	}
L66:
	;
	v179 = base.B2i32(v111 < v58)
	goto L43
}
func F_internal_citext_pattern_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	v7 = int32(1)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v11 = v9 & v7
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_pfree(m, v42)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L16
	} else {
		goto L56
	}
L2:
	;
	v12 = v7
	goto L4
L3:
	;
	v12 = int32(4)
	goto L4
L4:
	;
	if v9 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v42 = F_str_tolower(m, l0+v12, v40, int32(100))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v19 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v30 = int32(1)
	if v11 != 0 {
		v40 = int32(base.Ui32(v9)>>(uint(v30)%32)) - v30
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v22 = int32(16)
	goto L11
L10:
	;
	v22 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v29 = int32(4)
	goto L14
L13:
	;
	v29 = v22
	goto L14
L14:
	;
	v40 = v29
	goto L5
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	return int32(0)
L17:
	;
	v46 = int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v50 = v48 & v46
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v46
	goto L20
L19:
	;
	v51 = int32(4)
	goto L20
L20:
	;
	if v48 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v81 = F_str_tolower(m, l1+v51, v79, int32(100))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L16
	} else {
		goto L32
	}
L22:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v58 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v69 = int32(1)
	if v50 != 0 {
		v79 = int32(base.Ui32(v48)>>(uint(v69)%32)) - v69
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v61 = int32(16)
	goto L27
L26:
	;
	v61 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v68 = int32(4)
	goto L30
L29:
	;
	v68 = v61
	goto L30
L30:
	;
	v79 = v68
	goto L21
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v83 = F_strlen(m, v42)
	mBase = m.M
	v84 = F_strlen(m, v81)
	mBase = m.M
	v85 = base.B2i32(v83 < v84)
	if v83 < v84 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v86 = v83
	goto L35
L34:
	;
	v86 = v84
	goto L35
L35:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v86) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v148 != 0 {
		v151 = v148
		goto L1
	} else {
		goto L54
	}
L37:
	;
	v148 = int32(0)
	goto L36
L38:
	;
	v122 = v117
	v123 = v118
	v124 = v119
	goto L48
L39:
	;
	if (v42|v81)&int32(3) != 0 {
		v117 = v42
		v118 = v81
		v119 = v86
		goto L38
	} else {
		goto L42
	}
L40:
	;
	v110 = v42
	v111 = v81
	v112 = v86
	goto L41
L41:
	;
	if v112 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L42:
	;
	v94 = v42
	v95 = v81
	v96 = v86
	goto L43
L43:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v99 != v100 {
		v117 = v94
		v118 = v95
		v119 = v96
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v110 = v105
	v111 = v103
	v112 = v107
	goto L41
L45:
	;
	v102 = int32(4)
	v103 = v95 + v102
	v105 = v94 + v102
	v107 = v96 - v102
	if base.Ui32(int32(3)) < base.Ui32(v107) {
		v94 = v105
		v95 = v103
		v96 = v107
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v117 = v110
	v118 = v111
	v119 = v112
	goto L38
L48:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v127 == v128 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v148 = v127 - v128
	goto L36
L50:
	;
	v130 = int32(1)
	v135 = v124 - v130
	if v135 != 0 {
		v122 = v122 + v130
		v123 = v123 + v130
		v124 = v135
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	goto L37
L54:
	;
	if v83 < v84 {
		v151 = int32(-1)
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v151 = base.B2i32(v84 < v83)
	goto L1
L56:
	;
	F_pfree(m, v81)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	return v151
}
func F_internal_flush_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v6) < base.Ui32(v7) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = l0 + v7
	v11 = l0 + v6
	goto L5
L3:
	;
	goto L4
L4:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	goto L1
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[0]))
	v19 = F_secure_write(m, v17, v11, v9-v11)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if base.Ui32(v72) < base.Ui32(v9) {
		v11 = v72
		goto L5
	} else {
		goto L22
	}
L8:
	;
	return int32(0)
L9:
	;
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[1]))
	if v26 == int32(27) {
		v72 = v11
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[2])) = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v68 + v19
	v72 = v11 + v19
	goto L7
L13:
	;
	if v26 == int32(6) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[2]))
	if v26 == v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v53
	v58 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[3])) = v58
	*(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[4])) = v58
	return int32(-1)
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_internal_flush_buffer[2])) = v26
	v38 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	if v38 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F_internal_flush_buffer_0), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_internal_flush_buffer_1), int32(1405), int32(_a_F_internal_flush_buffer_2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	goto L6
}
func F_internal_get_result_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v275 int64
	_ = v275
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v685 int32
	_ = v685
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v977 int32
	_ = v977
	var v1001 int32
	_ = v1001
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v33 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L5
	} else {
		goto L282
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v947
	v977 = v953
	goto L1
L3:
	;
	v942 = int32(0)
	v943 = int32(3)
	if l4 == v942 {
		v977 = v943
		goto L1
	} else {
		goto L281
	}
L4:
	;
	if v40 <= int32(3830) {
		goto L251
	} else {
		goto L252
	}
L5:
	;
	return int32(0)
L6:
	;
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v39 = v37 + v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	v41 = F_build_function_result_tupdesc_t(m, v33)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L5
	} else {
		goto L245
	}
L10:
	;
	if v41 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v40
	goto L14
L13:
	;
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v46 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v812 != int32(2249) {
		goto L240
	} else {
		goto L241
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+128))
	v52 = v41 + v46<<(uint(int32(3))%32)
	v53 = int32(0)
	if v46 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v242&int32(1) == int32(0) {
		goto L15
	} else {
		goto L78
	}
L18:
	;
	v63 = v53
	v64 = v53
	v68 = v6
	v74 = v6
	v75 = v6
	v76 = v6
	v77 = v6
	v78 = v6
	v81 = v6
	v82 = v6
	v83 = v6
	goto L21
L19:
	;
	v181 = v53
	v182 = v53
	v192 = v6
	v193 = v6
	v194 = v6
	v195 = v6
	v196 = v6
	v199 = v6
	v200 = v6
	v201 = v6
	goto L20
L20:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v52+v182*int32(100))+96))
	if v208 <= int32(3830) {
		goto L63
	} else {
		goto L64
	}
L21:
	;
	v89 = v52 + v64*int32(100)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+96))
	if v90 <= int32(3830) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	if v46&int32(1) == int32(0) {
		v242 = v163
		v253 = v164
		v254 = v165
		v255 = v166
		v256 = v167
		v257 = v168
		v260 = v169
		v261 = v170
		v262 = v171
		goto L17
	} else {
		goto L60
	}
L23:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v89)+196))
	if v131 <= int32(3830) {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	v122 = int32(1)
	v123 = v113
	v124 = v114
	v125 = v115
	v126 = v116
	v127 = v117
	v128 = v118
	v129 = v119
	v130 = v120
	goto L23
L25:
	;
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v78
	v118 = v81
	v119 = v82
	v120 = int32(1)
	goto L24
L26:
	;
	switch v90 - int32(2277) {
	case 0:
		goto L25
	case 1, 2, 3, 4, 5:
		v122 = v63
		v123 = v74
		v124 = v75
		v125 = v76
		v126 = v77
		v127 = v78
		v128 = v81
		v129 = v82
		v130 = v83
		goto L23
	case 6:
		goto L29
	default:
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	switch v90 - int32(_a_F_internal_get_result_type_0) {
	case 0, 2:
		goto L36
	case 1:
		goto L35
	case 3:
		goto L34
	default:
		goto L37
	}
L29:
	;
	v113 = int32(1)
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v78
	v118 = v81
	v119 = v82
	v120 = v83
	goto L24
L30:
	;
	if v90 == int32(2776) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v90 != int32(3500) {
		v122 = v63
		v123 = v74
		v124 = v75
		v125 = v76
		v126 = v77
		v127 = v78
		v128 = v81
		v129 = v82
		v130 = v83
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v78
	v118 = v81
	v119 = int32(1)
	v120 = v83
	goto L24
L34:
	;
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v78
	v118 = int32(1)
	v119 = v82
	v120 = v83
	goto L24
L35:
	;
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = int32(1)
	v118 = v81
	v119 = v82
	v120 = v83
	goto L24
L36:
	;
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = int32(1)
	v117 = v78
	v118 = v81
	v119 = v82
	v120 = v83
	goto L24
L37:
	;
	switch v90 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L38
	case 1:
		goto L33
	default:
		goto L39
	}
L38:
	;
	v113 = v74
	v114 = v75
	v115 = int32(1)
	v116 = v77
	v117 = v78
	v118 = v81
	v119 = v82
	v120 = v83
	goto L24
L39:
	;
	if v90 != int32(3831) {
		v122 = v63
		v123 = v74
		v124 = v75
		v125 = v76
		v126 = v77
		v127 = v78
		v128 = v81
		v129 = v82
		v130 = v83
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v113 = v74
	v114 = int32(1)
	v115 = v76
	v116 = v77
	v117 = v78
	v118 = v81
	v119 = v82
	v120 = v83
	goto L24
L41:
	;
	v172 = int32(2)
	v173 = v64 + v172
	v175 = v68 + v172
	if v175 != v46&int32(2147483646) {
		v63 = v163
		v64 = v173
		v68 = v175
		v74 = v164
		v75 = v165
		v76 = v166
		v77 = v167
		v78 = v168
		v81 = v169
		v82 = v170
		v83 = v171
		goto L21
	} else {
		goto L59
	}
L42:
	;
	v163 = int32(1)
	v164 = v154
	v165 = v155
	v166 = v156
	v167 = v157
	v168 = v158
	v169 = v159
	v170 = v160
	v171 = v161
	goto L41
L43:
	;
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	v161 = int32(1)
	goto L42
L44:
	;
	switch v131 - int32(2277) {
	case 0:
		goto L43
	case 1, 2, 3, 4, 5:
		v163 = v122
		v164 = v123
		v165 = v124
		v166 = v125
		v167 = v126
		v168 = v127
		v169 = v128
		v170 = v129
		v171 = v130
		goto L41
	case 6:
		goto L47
	default:
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	switch v131 - int32(_a_F_internal_get_result_type_0) {
	case 0, 2:
		goto L52
	case 1:
		goto L53
	case 3:
		goto L54
	default:
		goto L55
	}
L47:
	;
	v154 = int32(1)
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	v161 = v130
	goto L42
L48:
	;
	if v131 == int32(2776) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v131 != int32(3500) {
		v163 = v122
		v164 = v123
		v165 = v124
		v166 = v125
		v167 = v126
		v168 = v127
		v169 = v128
		v170 = v129
		v171 = v130
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v154 = v123
	v155 = v124
	v156 = int32(1)
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	v161 = v130
	goto L42
L52:
	;
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = int32(1)
	v158 = v127
	v159 = v128
	v160 = v129
	v161 = v130
	goto L42
L53:
	;
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = int32(1)
	v159 = v128
	v160 = v129
	v161 = v130
	goto L42
L54:
	;
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = int32(1)
	v160 = v129
	v161 = v130
	goto L42
L55:
	;
	switch v131 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L51
	case 1:
		goto L56
	default:
		goto L57
	}
L56:
	;
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = int32(1)
	v161 = v130
	goto L42
L57:
	;
	if v131 != int32(3831) {
		v163 = v122
		v164 = v123
		v165 = v124
		v166 = v125
		v167 = v126
		v168 = v127
		v169 = v128
		v170 = v129
		v171 = v130
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v154 = v123
	v155 = int32(1)
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	v161 = v130
	goto L42
L59:
	;
	goto L22
L60:
	;
	v181 = v163
	v182 = v173
	v192 = v164
	v193 = v165
	v194 = v166
	v195 = v167
	v196 = v168
	v199 = v169
	v200 = v170
	v201 = v171
	goto L20
L61:
	;
	v242 = int32(1)
	v253 = v237
	v254 = v193
	v255 = v194
	v256 = v195
	v257 = v196
	v260 = v199
	v261 = v200
	v262 = v238
	goto L17
L62:
	;
	v237 = v192
	v238 = int32(1)
	goto L61
L63:
	;
	switch v208 - int32(2277) {
	case 0:
		goto L62
	case 1, 2, 3, 4, 5:
		v242 = v181
		v253 = v192
		v254 = v193
		v255 = v194
		v256 = v195
		v257 = v196
		v260 = v199
		v261 = v200
		v262 = v201
		goto L17
	case 6:
		goto L66
	default:
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	switch v208 - int32(_a_F_internal_get_result_type_0) {
	case 0, 2:
		goto L71
	case 1:
		goto L72
	case 3:
		goto L73
	default:
		goto L74
	}
L66:
	;
	v237 = int32(1)
	v238 = v201
	goto L61
L67:
	;
	if v208 == int32(2776) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if v208 != int32(3500) {
		v242 = v181
		v253 = v192
		v254 = v193
		v255 = v194
		v256 = v195
		v257 = v196
		v260 = v199
		v261 = v200
		v262 = v201
		goto L17
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v234 = int32(1)
	v242 = v234
	v253 = v192
	v254 = v193
	v255 = v234
	v256 = v195
	v257 = v196
	v260 = v199
	v261 = v200
	v262 = v201
	goto L17
L71:
	;
	v232 = int32(1)
	v242 = v232
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v232
	v257 = v196
	v260 = v199
	v261 = v200
	v262 = v201
	goto L17
L72:
	;
	v230 = int32(1)
	v242 = v230
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v195
	v257 = v230
	v260 = v199
	v261 = v200
	v262 = v201
	goto L17
L73:
	;
	v228 = int32(1)
	v242 = v228
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v195
	v257 = v196
	v260 = v228
	v261 = v200
	v262 = v201
	goto L17
L74:
	;
	switch v208 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L70
	case 1:
		goto L75
	default:
		goto L76
	}
L75:
	;
	v226 = int32(1)
	v242 = v226
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v195
	v257 = v196
	v260 = v199
	v261 = v226
	v262 = v201
	goto L17
L76:
	;
	if v208 != int32(3831) {
		v242 = v181
		v253 = v192
		v254 = v193
		v255 = v194
		v256 = v195
		v257 = v196
		v260 = v199
		v261 = v200
		v262 = v201
		goto L17
	} else {
		goto L77
	}
L77:
	;
	v224 = int32(1)
	v242 = v224
	v253 = v192
	v254 = v224
	v255 = v194
	v256 = v195
	v257 = v196
	v260 = v199
	v261 = v200
	v262 = v201
	goto L17
L78:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v272 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v272
	v275 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+52)) = v275
	*(*int64)(unsafe.Add(mBase, uint32(v29)+36)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v272
	if v272 < v49 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v288 = int32(0)
	v294 = v288
	v297 = v272
	v298 = v272
	v299 = v272
	v300 = v288
	v301 = v272
	v303 = v288
	v315 = v6
	v316 = v6
	goto L83
L81:
	;
	v389 = v272
	v390 = v272
	v391 = v272
	v393 = v272
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v390
	if base.B2i32(v390 == int32(0))&v253 != 0 {
		goto L126
	} else {
		goto L127
	}
L83:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(136)+v294<<(uint(int32(2))%32))))
	if v320 <= int32(3830) {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v368
	v389 = v366
	v390 = v367
	v391 = v368
	v393 = v370
	goto L82
L85:
	;
	v375 = v294 + int32(1)
	if v375 != v49 {
		v294 = v375
		v297 = v366
		v298 = v367
		v299 = v368
		v300 = v369
		v301 = v370
		v303 = v371
		v315 = v372
		v316 = v373
		goto L83
	} else {
		goto L125
	}
L86:
	;
	if v316 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L122
	}
L87:
	;
	if v303 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L119
	}
L88:
	;
	if v297 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L116
	}
L89:
	;
	if v301 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L113
	}
L90:
	;
	if v300 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L110
	}
L91:
	;
	if v299 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L107
	}
L92:
	;
	switch v320 - int32(2277) {
	case 0:
		goto L91
	case 1, 2, 3, 4, 5:
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	case 6:
		goto L95
	default:
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	switch v320 - int32(_a_F_internal_get_result_type_0) {
	case 0, 2:
		goto L89
	case 1:
		goto L88
	case 3:
		goto L87
	default:
		goto L102
	}
L95:
	;
	if v298 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L99
	}
L96:
	;
	if v320 == int32(2776) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v320 != int32(3500) {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v329 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	if v329 == int32(0) {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v366 = v297
	v367 = v329
	v368 = v299
	v369 = v300
	v370 = v301
	v371 = v303
	v372 = v315
	v373 = v316
	goto L85
L102:
	;
	switch v320 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L90
	case 1:
		goto L86
	default:
		goto L103
	}
L103:
	;
	if base.B2i32(v320 != int32(3831))|v315 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v303
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L104
	}
L104:
	;
	v340 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	if v340 == int32(0) {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v301
	v371 = v303
	v372 = v340
	v373 = v316
	goto L85
L107:
	;
	v344 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	if v344 == int32(0) {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v366 = v297
	v367 = v298
	v368 = v344
	v369 = v300
	v370 = v301
	v371 = v303
	v372 = v315
	v373 = v316
	goto L85
L110:
	;
	v348 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	if v348 == int32(0) {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v348
	v370 = v301
	v371 = v303
	v372 = v315
	v373 = v316
	goto L85
L113:
	;
	v352 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	if v352 == int32(0) {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v352
	v371 = v303
	v372 = v315
	v373 = v316
	goto L85
L116:
	;
	v356 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if v356 == int32(0) {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	v366 = v356
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v301
	v371 = v303
	v372 = v315
	v373 = v316
	goto L85
L119:
	;
	v360 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	if v360 != 0 {
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v301
		v371 = v360
		v372 = v315
		v373 = v316
		goto L85
	} else {
		goto L121
	}
L121:
	;
	goto L3
L122:
	;
	v362 = F_get_call_expr_argtype(m, l1, v294)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	if v362 == int32(0) {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v301
	v371 = v303
	v372 = v315
	v373 = v362
	goto L85
L125:
	;
	goto L84
L126:
	;
	F_resolve_anyelement_from_others(m, v27+int32(-16))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	v419 = v391
	goto L128
L128:
	;
	if base.B2i32(v419 == int32(0))&v262 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v419 = v418
	goto L128
L130:
	;
	F_resolve_anyarray_from_others(m, v27+int32(-16))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	if base.B2i32(v427 == int32(0))&v254 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	F_resolve_anyrange_from_others(m, v27+int32(-16))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L5
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if base.B2i32(v435 == int32(0))&v255 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L136
L138:
	;
	F_resolve_anymultirange_from_others(m, v27+int32(-16))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.B2i32(v393 == int32(0))&v256 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L140
L142:
	;
	F_resolve_anyelement_from_others(m, v27+int32(-32))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	v451 = v389
	goto L144
L144:
	;
	if base.B2i32(v451 == int32(0))&v257 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v451 = v450
	goto L144
L146:
	;
	F_resolve_anyarray_from_others(m, v27+int32(-32))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if base.B2i32(v459 == int32(0))&v260 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	F_resolve_anyrange_from_others(m, v27+int32(-32))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if base.B2i32(v467 == int32(0))&v261 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	F_resolve_anymultirange_from_others(m, v27+int32(-32))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L5
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v475 != 0 {
		v478 = v475
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L156
L158:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v483 != 0 {
		v486 = v483
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v479 = F_get_typcollation(m, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L162
	}
L160:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v476 != 0 {
		v478 = v476
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v482 = int32(0)
	goto L158
L162:
	;
	v482 = v479
	goto L158
L163:
	;
	v491 = int32(0)
	if v482|v490 == v491 {
		v532 = v491
		v533 = v491
		goto L168
	} else {
		goto L169
	}
L164:
	;
	v487 = F_get_typcollation(m, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L5
	} else {
		goto L167
	}
L165:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v484 != 0 {
		v486 = v484
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v490 = int32(0)
	goto L163
L167:
	;
	v490 = v487
	goto L163
L168:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v544 = int32(0)
	goto L189
L169:
	;
	v496 = int32(0)
	if l1 == v496 {
		v524 = v496
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v524 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L171:
	;
	goto L170
L172:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v503 = v501 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v503) {
		v524 = v496
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v507 = int32(1) << (uint(v503) % 32)
	if v507&int32(3904) == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v519+l1)))
	v524 = v521
	goto L171
L175:
	;
	if v507&int32(5) != 0 {
		v519 = int32(16)
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v519 = int32(24)
	goto L174
L178:
	;
	if v503 != int32(30) {
		v524 = v496
		goto L171
	} else {
		goto L179
	}
L179:
	;
	v519 = int32(12)
	goto L174
L180:
	;
	v532 = v482
	v533 = v490
	goto L168
L181:
	;
	goto L182
L182:
	;
	if v482 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v528 = v524
	goto L185
L184:
	;
	v528 = int32(0)
	goto L185
L185:
	;
	if v490 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v530 = v524
	goto L188
L187:
	;
	v530 = int32(0)
	goto L188
L188:
	;
	v532 = v528
	v533 = v530
	goto L168
L189:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v573 = v41 + v567<<(uint(int32(3))%32) + v544*int32(100)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v573)+96))
	if v576 <= int32(3830) {
		goto L199
	} else {
		goto L200
	}
L190:
	;
	v699 = int32(0)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v699 < v708 {
		goto L222
	} else {
		goto L223
	}
L191:
	;
	v697 = v544 + int32(1)
	if v697 != v46 {
		v544 = v697
		goto L189
	} else {
		goto L220
	}
L192:
	;
	F_TupleDescInitEntry(m, v41, base.I32_extend16_s(v544+int32(1)), v573+int32(32), v539, int32(-1), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L5
	} else {
		goto L219
	}
L193:
	;
	F_TupleDescInitEntry(m, v41, base.I32_extend16_s(v544+int32(1)), v573+int32(32), v538, int32(-1), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L5
	} else {
		goto L218
	}
L194:
	;
	v662 = base.I32_extend16_s(v544 + int32(1))
	F_TupleDescInitEntry(m, v41, v662, v573+int32(32), v537, int32(-1), int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L5
	} else {
		goto L216
	}
L195:
	;
	v645 = base.I32_extend16_s(v544 + int32(1))
	F_TupleDescInitEntry(m, v41, v645, v573+int32(32), v483, int32(-1), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L5
	} else {
		goto L214
	}
L196:
	;
	F_TupleDescInitEntry(m, v41, base.I32_extend16_s(v544+int32(1)), v573+int32(32), v536, int32(-1), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L5
	} else {
		goto L213
	}
L197:
	;
	if v576 != int32(3831) {
		goto L191
	} else {
		goto L211
	}
L198:
	;
	v608 = base.I32_extend16_s(v544 + int32(1))
	F_TupleDescInitEntry(m, v41, v608, v573+int32(32), v534, int32(-1), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L5
	} else {
		goto L209
	}
L199:
	;
	switch v576 - int32(2277) {
	case 0:
		goto L198
	case 1, 2, 3, 4, 5:
		goto L191
	case 6:
		goto L202
	default:
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	switch v576 - int32(_a_F_internal_get_result_type_0) {
	case 0, 2:
		goto L195
	case 1:
		goto L194
	case 3:
		goto L193
	default:
		goto L208
	}
L202:
	;
	v587 = base.I32_extend16_s(v544 + int32(1))
	F_TupleDescInitEntry(m, v41, v587, v573+int32(32), v475, int32(-1), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L5
	} else {
		goto L206
	}
L203:
	;
	if v576 == int32(2776) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	if v576 != int32(3500) {
		goto L191
	} else {
		goto L205
	}
L205:
	;
	goto L202
L206:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v594<<(uint(int32(3))%32)+v587*int32(100))+24)) = v532
	goto L207
L207:
	;
	goto L191
L208:
	;
	switch v576 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L196
	case 1:
		goto L192
	default:
		goto L197
	}
L209:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v615<<(uint(int32(3))%32)+v608*int32(100))+24)) = v532
	goto L210
L210:
	;
	goto L191
L211:
	;
	F_TupleDescInitEntry(m, v41, base.I32_extend16_s(v544+int32(1)), v573+int32(32), v535, int32(-1), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	goto L191
L213:
	;
	goto L191
L214:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v652<<(uint(int32(3))%32)+v645*int32(100))+24)) = v533
	goto L215
L215:
	;
	goto L191
L216:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v669<<(uint(int32(3))%32)+v662*int32(100))+24)) = v533
	goto L217
L217:
	;
	goto L191
L218:
	;
	goto L191
L219:
	;
	goto L191
L220:
	;
	goto L190
L221:
	;
	goto L15
L222:
	;
	v712 = v41 + int32(28)
	v719 = v699
	v720 = v708
	v722 = v699
	goto L226
L223:
	;
	v776 = v699
	v783 = v708
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v783
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v776
	goto L221
L225:
	;
	v776 = v770
	v783 = v749
	goto L224
L226:
	;
	v728 = v712 + v708<<(uint(int32(3))%32) + v719*int32(100)
	v731 = v712 + v719<<(uint(int32(3))%32)
	if v708 != v720 {
		v749 = v720
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v770 = v708
	goto L225
L228:
	;
	v750 = int32(*(*int16)(unsafe.Add(mBase, uint32(v731)+2)))
	if v750 <= int32(0) {
		v770 = v719
		goto L225
	} else {
		goto L236
	}
L229:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+7)))
	if v733 != int32(118) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v749 = v719
	goto L228
L231:
	;
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+4)))
	if v736 != int32(1) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+6)))
	if v739&int32(6) != 0 {
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v742 = int32(*(*int16)(unsafe.Add(mBase, uint32(v731)+2)))
	if v742 <= int32(0) {
		goto L230
	} else {
		goto L234
	}
L234:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+90)))
	if v745 != int32(118) {
		v749 = v708
		goto L228
	} else {
		goto L235
	}
L235:
	;
	goto L230
L236:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728)+90)))
	if v753 == int32(118) {
		v770 = v719
		goto L225
	} else {
		goto L237
	}
L237:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+5)))
	v762 = (v722 + v756 - int32(1)) & (int32(0) - v756)
	if int32(_a_F_internal_get_result_type_2) < v762 {
		v770 = v719
		goto L225
	} else {
		goto L238
	}
L238:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v731))) = uint16(v762)
	v768 = v719 + int32(1)
	if v768 != v708 {
		v719 = v768
		v720 = v749
		v722 = v762 + v750
		goto L226
	} else {
		goto L239
	}
L239:
	;
	goto L227
L240:
	;
	v820 = int32(1)
	if l4 != 0 {
		v947 = v41
		v953 = v820
		goto L2
	} else {
		goto L244
	}
L241:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if int32(0) <= v815 {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	F_assign_record_type_typmod(m, v41)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L243
	}
L243:
	;
	goto L240
L244:
	;
	v977 = v820
	goto L1
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = l0
	F_errmsg_internal(m, int32(_a_F_internal_get_result_type_3), v29)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_internal_get_result_type_4), int32(448), int32(_a_F_internal_get_result_type_5))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L5
	} else {
		goto L276
	}
L249:
	;
	if l3 != 0 {
		goto L259
	} else {
		goto L260
	}
L250:
	;
	v857 = F_exprType(m, l1)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L5
	} else {
		goto L257
	}
L251:
	;
	switch v40 - int32(2277) {
	case 0, 6:
		goto L250
	case 1, 2, 3, 4, 5:
		v861 = v40
		goto L249
	default:
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	if base.B2i32(base.Ui32(v40-int32(_a_F_internal_get_result_type_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v40-int32(_a_F_internal_get_result_type_1)) < base.Ui32(int32(2)))|base.B2i32(v40 == int32(3831)) != 0 {
		goto L250
	} else {
		goto L256
	}
L254:
	;
	if base.B2i32(v40 == int32(2776))|base.B2i32(v40 == int32(3500)) != 0 {
		goto L250
	} else {
		goto L255
	}
L255:
	;
	v861 = v40
	goto L249
L256:
	;
	v861 = v40
	goto L249
L257:
	;
	if v857 == int32(0) {
		goto L248
	} else {
		goto L258
	}
L258:
	;
	v861 = v857
	goto L249
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v861
	goto L261
L260:
	;
	goto L261
L261:
	;
	if l4 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L264
L263:
	;
	goto L264
L264:
	;
	v867 = F_get_type_func_class(m, v861, v27+int32(-16))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L267
	}
L265:
	;
	v877 = int32(3)
	if l2 == int32(0) {
		v977 = v877
		goto L1
	} else {
		goto L270
	}
L266:
	;
	if l4 == int32(0) {
		v977 = v867
		goto L1
	} else {
		goto L268
	}
L267:
	;
	switch v867 - int32(1) {
	case 0, 1:
		goto L266
	case 2:
		goto L265
	default:
		v977 = v867
		goto L1
	}
L268:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v875 = F_lookup_rowtype_tupdesc_copy(m, v873, int32(-1))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L269
	}
L269:
	;
	v947 = v875
	v953 = v867
	goto L2
L270:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v880 != int32(389) {
		v977 = v877
		goto L1
	} else {
		goto L271
	}
L271:
	;
	v883 = int32(1)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v886 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v887 = v883
	goto L274
L273:
	;
	v887 = int32(3)
	goto L274
L274:
	;
	v888 = int32(0)
	if base.B2i32(l4 == v888)|base.B2i32(v886 == v888) != 0 {
		v977 = v887
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v947 = v886
	v953 = v883
	goto L2
L276:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	v900 = F_format_type_be(m, v40)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v39 + int32(4)
	F_errmsg(m, int32(_a_F_internal_get_result_type_6), v27+int32(-48))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_internal_get_result_type_4), int32(500), int32(_a_F_internal_get_result_type_5))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L5
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	v947 = v942
	v953 = v943
	goto L2
L282:
	;
	m.G0 = v29 - int32(-64)
	return v977
}
func F_internal_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_internal_in_0), int32(373), int32(_a_F_internal_in_1), int32(_a_F_internal_in_2), int32(_a_F_internal_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
