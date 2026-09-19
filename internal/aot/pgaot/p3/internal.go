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
	if v89 == int32(3) {
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
											F_ReportApplyConflict(m, v13, l1, int32(15), int32(4), l2, int32(0), v65)
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
								F_ReportApplyConflict(m, v13, l1, int32(15), int32(5), l2, int32(0), v89)
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
								F_ReportApplyConflict(m, v13, l1, int32(15), int32(5), l2, int32(0), v89)
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
											F_ReportApplyConflict(m, v13, l1, int32(15), int32(4), l2, int32(0), v65)
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
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
							F_errfinish(m, int32(_a_F_attach_internal_1), int32(1364), int32(_a_F_attach_internal_2))
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	v10 = int32(1)
	v11 = l0 + v10
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = v14 & v10
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v11
	goto L3
L2:
	;
	v17 = l0 + int32(4)
	goto L3
L3:
	;
	if v14 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v51 = v44
	goto L15
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v23 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16 != 0 {
		v44 = int32(base.Ui32(v14)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v26 = int32(16)
	goto L10
L9:
	;
	v26 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = int32(4)
	goto L13
L12:
	;
	v33 = v26
	goto L13
L13:
	;
	v44 = v33
	goto L4
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v51 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v67 = int32(1)
	v68 = l1 + v67
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v73 = v71 & v67
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v66 = v44 & (v44 >> (uint(int32(31)) % 32))
	goto L17
L19:
	;
	goto L20
L20:
	;
	v60 = v51 - int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v60))))
	if v62 == int32(32) {
		v51 = v60
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v66 = v51
	goto L17
L22:
	;
	v74 = v68
	goto L24
L23:
	;
	v74 = l1 + int32(4)
	goto L24
L24:
	;
	if v71 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v108 = v101
	goto L36
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v80 == int32(18) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v91 = int32(1)
	if v73 != 0 {
		v101 = int32(base.Ui32(v71)>>(uint(v91)%32)) - v91
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v83 = int32(16)
	goto L31
L30:
	;
	v83 = int32(0)
	goto L31
L31:
	;
	if base.Ui32((v80-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v90 = int32(4)
	goto L34
L33:
	;
	v90 = v83
	goto L34
L34:
	;
	v101 = v90
	goto L25
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	if v108 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v124 = int32(1)
	if v14&v124 != 0 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	goto L37
L39:
	;
	v122 = v101 & (v101 >> (uint(int32(31)) % 32))
	goto L38
L40:
	;
	goto L41
L41:
	;
	v117 = v108 - int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v117))))
	if v119 == int32(32) {
		v108 = v117
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v122 = v108
	goto L38
L43:
	;
	return v202
L44:
	;
	v128 = v124
	goto L46
L45:
	;
	v128 = int32(4)
	goto L46
L46:
	;
	v129 = l0 + v128
	v130 = int32(1)
	if v71&v130 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v134 = v130
	goto L49
L48:
	;
	v134 = int32(4)
	goto L49
L49:
	;
	v135 = l1 + v134
	v136 = base.B2i32(v66 < v122)
	if v66 < v122 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v137 = v66
	goto L52
L51:
	;
	v137 = v122
	goto L52
L52:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v137) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v199 != 0 {
		v202 = v199
		goto L43
	} else {
		goto L71
	}
L54:
	;
	v199 = int32(0)
	goto L53
L55:
	;
	v173 = v168
	v174 = v169
	v175 = v170
	goto L65
L56:
	;
	if (v129|v135)&int32(3) != 0 {
		v168 = v129
		v169 = v135
		v170 = v137
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v161 = v129
	v162 = v135
	v163 = v137
	goto L58
L58:
	;
	if v163 == int32(0) {
		goto L54
	} else {
		goto L64
	}
L59:
	;
	v145 = v129
	v146 = v135
	v147 = v137
	goto L60
L60:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v150 != v151 {
		v168 = v145
		v169 = v146
		v170 = v147
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v161 = v156
	v162 = v154
	v163 = v158
	goto L58
L62:
	;
	v153 = int32(4)
	v154 = v146 + v153
	v156 = v145 + v153
	v158 = v147 - v153
	if base.Ui32(int32(3)) < base.Ui32(v158) {
		v145 = v156
		v146 = v154
		v147 = v158
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v168 = v161
	v169 = v162
	v170 = v163
	goto L55
L65:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v178 == v179 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v199 = v178 - v179
	goto L53
L67:
	;
	v181 = int32(1)
	v186 = v175 - v181
	if v186 != 0 {
		v173 = v173 + v181
		v174 = v174 + v181
		v175 = v186
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L54
L71:
	;
	if v66 < v122 {
		v202 = int32(-1)
		goto L43
	} else {
		goto L72
	}
L72:
	;
	v202 = base.B2i32(v122 < v66)
	goto L43
}
func F_internal_citext_pattern_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	v7 = int32(1)
	v8 = l0 + v7
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v13 = v11 & v7
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_pfree(m, v43)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L16
	} else {
		goto L56
	}
L2:
	;
	v14 = v8
	goto L4
L3:
	;
	v14 = l0 + int32(4)
	goto L4
L4:
	;
	if v11 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v43 = F_str_tolower(m, v14, v41, int32(100))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v20 == int32(18) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v31 = int32(1)
	if v13 != 0 {
		v41 = int32(base.Ui32(v11)>>(uint(v31)%32)) - v31
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v23 = int32(16)
	goto L11
L10:
	;
	v23 = int32(0)
	goto L11
L11:
	;
	if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v30 = int32(4)
	goto L14
L13:
	;
	v30 = v23
	goto L14
L14:
	;
	v41 = v30
	goto L5
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	return int32(0)
L17:
	;
	v47 = int32(1)
	v48 = l1 + v47
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v53 = v51 & v47
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v54 = v48
	goto L20
L19:
	;
	v54 = l1 + int32(4)
	goto L20
L20:
	;
	if v51 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v83 = F_str_tolower(m, v54, v81, int32(100))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L32
	}
L22:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v60 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v71 = int32(1)
	if v53 != 0 {
		v81 = int32(base.Ui32(v51)>>(uint(v71)%32)) - v71
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v63 = int32(16)
	goto L27
L26:
	;
	v63 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v60-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v70 = int32(4)
	goto L30
L29:
	;
	v70 = v63
	goto L30
L30:
	;
	v81 = v70
	goto L21
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v85 = F_strlen(m, v43)
	mBase = m.M
	v86 = F_strlen(m, v83)
	mBase = m.M
	v87 = base.B2i32(v85 < v86)
	if v85 < v86 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v88 = v85
	goto L35
L34:
	;
	v88 = v86
	goto L35
L35:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v88) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v150 != 0 {
		v153 = v150
		goto L1
	} else {
		goto L54
	}
L37:
	;
	v150 = int32(0)
	goto L36
L38:
	;
	v124 = v119
	v125 = v120
	v126 = v121
	goto L48
L39:
	;
	if (v43|v83)&int32(3) != 0 {
		v119 = v43
		v120 = v83
		v121 = v88
		goto L38
	} else {
		goto L42
	}
L40:
	;
	v112 = v43
	v113 = v83
	v114 = v88
	goto L41
L41:
	;
	if v114 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L42:
	;
	v96 = v43
	v97 = v83
	v98 = v88
	goto L43
L43:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v101 != v102 {
		v119 = v96
		v120 = v97
		v121 = v98
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v112 = v107
	v113 = v105
	v114 = v109
	goto L41
L45:
	;
	v104 = int32(4)
	v105 = v97 + v104
	v107 = v96 + v104
	v109 = v98 - v104
	if base.Ui32(int32(3)) < base.Ui32(v109) {
		v96 = v107
		v97 = v105
		v98 = v109
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v119 = v112
	v120 = v113
	v121 = v114
	goto L38
L48:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v129 == v130 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v150 = v129 - v130
	goto L36
L50:
	;
	v132 = int32(1)
	v137 = v126 - v132
	if v137 != 0 {
		v124 = v124 + v132
		v125 = v125 + v132
		v126 = v137
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
	if v85 < v86 {
		v153 = int32(-1)
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v153 = base.B2i32(v86 < v85)
	goto L1
L56:
	;
	F_pfree(m, v83)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	return v153
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
	var v32 int32
	_ = v32
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v112 int32
	_ = v112
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
	var v121 int32
	_ = v121
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
	var v153 int32
	_ = v153
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
	var v162 int32
	_ = v162
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
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v274 int64
	_ = v274
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
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
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
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
	var v543 int32
	_ = v543
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v889 int32
	_ = v889
	var v913 int32
	_ = v913
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v32 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L5
	} else {
		goto L263
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v859
	v889 = v865
	goto L1
L3:
	;
	v854 = int32(0)
	v855 = int32(3)
	if l4 == v854 {
		v889 = v855
		goto L1
	} else {
		goto L262
	}
L4:
	;
	if v39 <= int32(3830) {
		goto L232
	} else {
		goto L233
	}
L5:
	;
	return int32(0)
L6:
	;
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = v36 + v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+108))
	v40 = F_build_function_result_tupdesc_t(m, v32)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v737 = m.ExcPending
	if v737 != 0 {
		goto L5
	} else {
		goto L226
	}
L10:
	;
	if v40 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v39
	goto L14
L13:
	;
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v45 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v724 != int32(2249) {
		goto L221
	} else {
		goto L222
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+128))
	v51 = v40 + v45<<(uint(int32(4))%32)
	v52 = int32(0)
	if v45 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v241&int32(1) == int32(0) {
		goto L15
	} else {
		goto L78
	}
L18:
	;
	v62 = v52
	v63 = v52
	v67 = v6
	v73 = v6
	v74 = v6
	v75 = v6
	v76 = v6
	v77 = v6
	v80 = v6
	v81 = v6
	v82 = v6
	goto L21
L19:
	;
	v180 = v52
	v181 = v52
	v191 = v6
	v192 = v6
	v193 = v6
	v194 = v6
	v195 = v6
	v198 = v6
	v199 = v6
	v200 = v6
	goto L20
L20:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v51+v181*int32(100))+88))
	if v207 <= int32(3830) {
		goto L63
	} else {
		goto L64
	}
L21:
	;
	v88 = v51 + v63*int32(100)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+88))
	if v89 <= int32(3830) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	if v45&int32(1) == int32(0) {
		v241 = v162
		v252 = v163
		v253 = v164
		v254 = v165
		v255 = v166
		v256 = v167
		v259 = v168
		v260 = v169
		v261 = v170
		goto L17
	} else {
		goto L60
	}
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v88)+188))
	if v130 <= int32(3830) {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	v121 = int32(1)
	v122 = v112
	v123 = v113
	v124 = v114
	v125 = v115
	v126 = v116
	v127 = v117
	v128 = v118
	v129 = v119
	goto L23
L25:
	;
	v112 = v73
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v80
	v118 = v81
	v119 = int32(1)
	goto L24
L26:
	;
	switch v89 - int32(2277) {
	case 0:
		goto L25
	case 1, 2, 3, 4, 5:
		v121 = v62
		v122 = v73
		v123 = v74
		v124 = v75
		v125 = v76
		v126 = v77
		v127 = v80
		v128 = v81
		v129 = v82
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
	switch v89 - int32(_a_F_internal_get_result_type_0) {
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
	v112 = int32(1)
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v80
	v118 = v81
	v119 = v82
	goto L24
L30:
	;
	if v89 == int32(2776) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v89 != int32(3500) {
		v121 = v62
		v122 = v73
		v123 = v74
		v124 = v75
		v125 = v76
		v126 = v77
		v127 = v80
		v128 = v81
		v129 = v82
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v112 = v73
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v80
	v118 = int32(1)
	v119 = v82
	goto L24
L34:
	;
	v112 = v73
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = int32(1)
	v118 = v81
	v119 = v82
	goto L24
L35:
	;
	v112 = v73
	v113 = v74
	v114 = v75
	v115 = v76
	v116 = int32(1)
	v117 = v80
	v118 = v81
	v119 = v82
	goto L24
L36:
	;
	v112 = v73
	v113 = v74
	v114 = v75
	v115 = int32(1)
	v116 = v77
	v117 = v80
	v118 = v81
	v119 = v82
	goto L24
L37:
	;
	switch v89 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L38
	case 1:
		goto L33
	default:
		goto L39
	}
L38:
	;
	v112 = v73
	v113 = v74
	v114 = int32(1)
	v115 = v76
	v116 = v77
	v117 = v80
	v118 = v81
	v119 = v82
	goto L24
L39:
	;
	if v89 != int32(3831) {
		v121 = v62
		v122 = v73
		v123 = v74
		v124 = v75
		v125 = v76
		v126 = v77
		v127 = v80
		v128 = v81
		v129 = v82
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v112 = v73
	v113 = int32(1)
	v114 = v75
	v115 = v76
	v116 = v77
	v117 = v80
	v118 = v81
	v119 = v82
	goto L24
L41:
	;
	v171 = int32(2)
	v172 = v63 + v171
	v174 = v67 + v171
	if v174 != v45&int32(2147483646) {
		v62 = v162
		v63 = v172
		v67 = v174
		v73 = v163
		v74 = v164
		v75 = v165
		v76 = v166
		v77 = v167
		v80 = v168
		v81 = v169
		v82 = v170
		goto L21
	} else {
		goto L59
	}
L42:
	;
	v162 = int32(1)
	v163 = v153
	v164 = v154
	v165 = v155
	v166 = v156
	v167 = v157
	v168 = v158
	v169 = v159
	v170 = v160
	goto L41
L43:
	;
	v153 = v122
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = int32(1)
	goto L42
L44:
	;
	switch v130 - int32(2277) {
	case 0:
		goto L43
	case 1, 2, 3, 4, 5:
		v162 = v121
		v163 = v122
		v164 = v123
		v165 = v124
		v166 = v125
		v167 = v126
		v168 = v127
		v169 = v128
		v170 = v129
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
	switch v130 - int32(_a_F_internal_get_result_type_0) {
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
	v153 = int32(1)
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	goto L42
L48:
	;
	if v130 == int32(2776) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	if v130 != int32(3500) {
		v162 = v121
		v163 = v122
		v164 = v123
		v165 = v124
		v166 = v125
		v167 = v126
		v168 = v127
		v169 = v128
		v170 = v129
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v153 = v122
	v154 = v123
	v155 = int32(1)
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	goto L42
L52:
	;
	v153 = v122
	v154 = v123
	v155 = v124
	v156 = int32(1)
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	goto L42
L53:
	;
	v153 = v122
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = int32(1)
	v158 = v127
	v159 = v128
	v160 = v129
	goto L42
L54:
	;
	v153 = v122
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = int32(1)
	v159 = v128
	v160 = v129
	goto L42
L55:
	;
	switch v130 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L51
	case 1:
		goto L56
	default:
		goto L57
	}
L56:
	;
	v153 = v122
	v154 = v123
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = int32(1)
	v160 = v129
	goto L42
L57:
	;
	if v130 != int32(3831) {
		v162 = v121
		v163 = v122
		v164 = v123
		v165 = v124
		v166 = v125
		v167 = v126
		v168 = v127
		v169 = v128
		v170 = v129
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v153 = v122
	v154 = int32(1)
	v155 = v124
	v156 = v125
	v157 = v126
	v158 = v127
	v159 = v128
	v160 = v129
	goto L42
L59:
	;
	goto L22
L60:
	;
	v180 = v162
	v181 = v172
	v191 = v163
	v192 = v164
	v193 = v165
	v194 = v166
	v195 = v167
	v198 = v168
	v199 = v169
	v200 = v170
	goto L20
L61:
	;
	v241 = int32(1)
	v252 = v236
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v195
	v259 = v198
	v260 = v199
	v261 = v237
	goto L17
L62:
	;
	v236 = v191
	v237 = int32(1)
	goto L61
L63:
	;
	switch v207 - int32(2277) {
	case 0:
		goto L62
	case 1, 2, 3, 4, 5:
		v241 = v180
		v252 = v191
		v253 = v192
		v254 = v193
		v255 = v194
		v256 = v195
		v259 = v198
		v260 = v199
		v261 = v200
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
	switch v207 - int32(_a_F_internal_get_result_type_0) {
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
	v236 = int32(1)
	v237 = v200
	goto L61
L67:
	;
	if v207 == int32(2776) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if v207 != int32(3500) {
		v241 = v180
		v252 = v191
		v253 = v192
		v254 = v193
		v255 = v194
		v256 = v195
		v259 = v198
		v260 = v199
		v261 = v200
		goto L17
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	v233 = int32(1)
	v241 = v233
	v252 = v191
	v253 = v192
	v254 = v233
	v255 = v194
	v256 = v195
	v259 = v198
	v260 = v199
	v261 = v200
	goto L17
L71:
	;
	v231 = int32(1)
	v241 = v231
	v252 = v191
	v253 = v192
	v254 = v193
	v255 = v231
	v256 = v195
	v259 = v198
	v260 = v199
	v261 = v200
	goto L17
L72:
	;
	v229 = int32(1)
	v241 = v229
	v252 = v191
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v229
	v259 = v198
	v260 = v199
	v261 = v200
	goto L17
L73:
	;
	v227 = int32(1)
	v241 = v227
	v252 = v191
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v195
	v259 = v227
	v260 = v199
	v261 = v200
	goto L17
L74:
	;
	switch v207 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L70
	case 1:
		goto L75
	default:
		goto L76
	}
L75:
	;
	v225 = int32(1)
	v241 = v225
	v252 = v191
	v253 = v192
	v254 = v193
	v255 = v194
	v256 = v195
	v259 = v198
	v260 = v225
	v261 = v200
	goto L17
L76:
	;
	if v207 != int32(3831) {
		v241 = v180
		v252 = v191
		v253 = v192
		v254 = v193
		v255 = v194
		v256 = v195
		v259 = v198
		v260 = v199
		v261 = v200
		goto L17
	} else {
		goto L77
	}
L77:
	;
	v223 = int32(1)
	v241 = v223
	v252 = v191
	v253 = v223
	v254 = v193
	v255 = v194
	v256 = v195
	v259 = v198
	v260 = v199
	v261 = v200
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
	v271 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v271
	v274 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+52)) = v274
	*(*int64)(unsafe.Add(mBase, uint32(v29)+36)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v271
	if v271 < v48 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v287 = int32(0)
	v293 = v287
	v296 = v271
	v297 = v271
	v298 = v271
	v299 = v287
	v300 = v271
	v302 = v287
	v314 = v6
	v315 = v6
	goto L83
L81:
	;
	v388 = v271
	v389 = v271
	v390 = v271
	v392 = v271
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v389
	if base.B2i32(v389 == int32(0))&v252 != 0 {
		goto L126
	} else {
		goto L127
	}
L83:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(136)+v293<<(uint(int32(2))%32))))
	if v319 <= int32(3830) {
		goto L92
	} else {
		goto L93
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v367
	v388 = v365
	v389 = v366
	v390 = v367
	v392 = v369
	goto L82
L85:
	;
	v374 = v293 + int32(1)
	if v374 != v48 {
		v293 = v374
		v296 = v365
		v297 = v366
		v298 = v367
		v299 = v368
		v300 = v369
		v302 = v370
		v314 = v371
		v315 = v372
		goto L83
	} else {
		goto L125
	}
L86:
	;
	if v315 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L122
	}
L87:
	;
	if v302 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L119
	}
L88:
	;
	if v296 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L116
	}
L89:
	;
	if v300 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L113
	}
L90:
	;
	if v299 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L110
	}
L91:
	;
	if v298 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L107
	}
L92:
	;
	switch v319 - int32(2277) {
	case 0:
		goto L91
	case 1, 2, 3, 4, 5:
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
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
	switch v319 - int32(_a_F_internal_get_result_type_0) {
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
	if v297 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L99
	}
L96:
	;
	if v319 == int32(2776) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	if v319 != int32(3500) {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	v328 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	if v328 == int32(0) {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v365 = v296
	v366 = v328
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v302
	v371 = v314
	v372 = v315
	goto L85
L102:
	;
	switch v319 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L90
	case 1:
		goto L86
	default:
		goto L103
	}
L103:
	;
	if base.B2i32(v319 != int32(3831))|v314 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v302
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L104
	}
L104:
	;
	v339 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	if v339 == int32(0) {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	v365 = v296
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v302
	v371 = v339
	v372 = v315
	goto L85
L107:
	;
	v343 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	if v343 == int32(0) {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v365 = v296
	v366 = v297
	v367 = v343
	v368 = v299
	v369 = v300
	v370 = v302
	v371 = v314
	v372 = v315
	goto L85
L110:
	;
	v347 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	if v347 == int32(0) {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	v365 = v296
	v366 = v297
	v367 = v298
	v368 = v347
	v369 = v300
	v370 = v302
	v371 = v314
	v372 = v315
	goto L85
L113:
	;
	v351 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	if v351 == int32(0) {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v365 = v296
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v351
	v370 = v302
	v371 = v314
	v372 = v315
	goto L85
L116:
	;
	v355 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	if v355 == int32(0) {
		goto L3
	} else {
		goto L118
	}
L118:
	;
	v365 = v355
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v302
	v371 = v314
	v372 = v315
	goto L85
L119:
	;
	v359 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	if v359 != 0 {
		v365 = v296
		v366 = v297
		v367 = v298
		v368 = v299
		v369 = v300
		v370 = v359
		v371 = v314
		v372 = v315
		goto L85
	} else {
		goto L121
	}
L121:
	;
	goto L3
L122:
	;
	v361 = F_get_call_expr_argtype(m, l1, v293)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	if v361 == int32(0) {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	v365 = v296
	v366 = v297
	v367 = v298
	v368 = v299
	v369 = v300
	v370 = v302
	v371 = v314
	v372 = v361
	goto L85
L125:
	;
	goto L84
L126:
	;
	F_resolve_anyelement_from_others(m, v27+int32(-16))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	v418 = v390
	goto L128
L128:
	;
	if base.B2i32(v418 == int32(0))&v261 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v418 = v417
	goto L128
L130:
	;
	F_resolve_anyarray_from_others(m, v27+int32(-16))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	if base.B2i32(v426 == int32(0))&v253 != 0 {
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
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if base.B2i32(v434 == int32(0))&v254 != 0 {
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
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.B2i32(v392 == int32(0))&v255 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	v450 = v388
	goto L144
L144:
	;
	if base.B2i32(v450 == int32(0))&v256 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v450 = v449
	goto L144
L146:
	;
	F_resolve_anyarray_from_others(m, v27+int32(-32))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	if base.B2i32(v458 == int32(0))&v259 != 0 {
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
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if base.B2i32(v466 == int32(0))&v260 != 0 {
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
	v473 = m.ExcPending
	if v473 != 0 {
		goto L5
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v474 != 0 {
		v477 = v474
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L156
L158:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v482 != 0 {
		v485 = v482
		goto L164
	} else {
		goto L165
	}
L159:
	;
	v478 = F_get_typcollation(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L162
	}
L160:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v475 != 0 {
		v477 = v475
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v481 = int32(0)
	goto L158
L162:
	;
	v481 = v478
	goto L158
L163:
	;
	v490 = int32(0)
	if v481|v489 == v490 {
		v531 = v490
		v532 = v490
		goto L168
	} else {
		goto L169
	}
L164:
	;
	v486 = F_get_typcollation(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L167
	}
L165:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v483 != 0 {
		v485 = v483
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v489 = int32(0)
	goto L163
L167:
	;
	v489 = v486
	goto L163
L168:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v543 = int32(0)
	goto L189
L169:
	;
	v495 = int32(0)
	if l1 == v495 {
		v523 = v495
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v523 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L171:
	;
	goto L170
L172:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v502 = v500 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v502) {
		v523 = v495
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v506 = int32(1) << (uint(v502) % 32)
	if v506&int32(3904) == int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v518+l1)))
	v523 = v520
	goto L171
L175:
	;
	if v506&int32(5) != 0 {
		v518 = int32(16)
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v518 = int32(24)
	goto L174
L178:
	;
	if v502 != int32(30) {
		v523 = v495
		goto L171
	} else {
		goto L179
	}
L179:
	;
	v518 = int32(12)
	goto L174
L180:
	;
	v531 = v481
	v532 = v489
	goto L168
L181:
	;
	goto L182
L182:
	;
	if v481 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v527 = v523
	goto L185
L184:
	;
	v527 = int32(0)
	goto L185
L185:
	;
	if v489 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v529 = v523
	goto L188
L187:
	;
	v529 = int32(0)
	goto L188
L188:
	;
	v531 = v527
	v532 = v529
	goto L168
L189:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v572 = v40 + v566<<(uint(int32(4))%32) + v543*int32(100)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v572)+88))
	if v575 <= int32(3830) {
		goto L199
	} else {
		goto L200
	}
L190:
	;
	goto L15
L191:
	;
	v696 = v543 + int32(1)
	if v696 != v45 {
		v543 = v696
		goto L189
	} else {
		goto L220
	}
L192:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v543+int32(1)), v572+int32(24), v538, int32(-1), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L219
	}
L193:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v543+int32(1)), v572+int32(24), v537, int32(-1), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L5
	} else {
		goto L218
	}
L194:
	;
	v661 = base.I32_extend16_s(v543 + int32(1))
	F_TupleDescInitEntry(m, v40, v661, v572+int32(24), v536, int32(-1), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L5
	} else {
		goto L216
	}
L195:
	;
	v644 = base.I32_extend16_s(v543 + int32(1))
	F_TupleDescInitEntry(m, v40, v644, v572+int32(24), v482, int32(-1), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L5
	} else {
		goto L214
	}
L196:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v543+int32(1)), v572+int32(24), v535, int32(-1), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L5
	} else {
		goto L213
	}
L197:
	;
	if v575 != int32(3831) {
		goto L191
	} else {
		goto L211
	}
L198:
	;
	v607 = base.I32_extend16_s(v543 + int32(1))
	F_TupleDescInitEntry(m, v40, v607, v572+int32(24), v533, int32(-1), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L5
	} else {
		goto L209
	}
L199:
	;
	switch v575 - int32(2277) {
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
	switch v575 - int32(_a_F_internal_get_result_type_0) {
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
	v586 = base.I32_extend16_s(v543 + int32(1))
	F_TupleDescInitEntry(m, v40, v586, v572+int32(24), v474, int32(-1), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L5
	} else {
		goto L206
	}
L203:
	;
	if v575 == int32(2776) {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	if v575 != int32(3500) {
		goto L191
	} else {
		goto L205
	}
L205:
	;
	goto L202
L206:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v593<<(uint(int32(4))%32)+v586*int32(100))+16)) = v531
	goto L207
L207:
	;
	goto L191
L208:
	;
	switch v575 - int32(_a_F_internal_get_result_type_1) {
	case 0:
		goto L196
	case 1:
		goto L192
	default:
		goto L197
	}
L209:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v614<<(uint(int32(4))%32)+v607*int32(100))+16)) = v531
	goto L210
L210:
	;
	goto L191
L211:
	;
	F_TupleDescInitEntry(m, v40, base.I32_extend16_s(v543+int32(1)), v572+int32(24), v534, int32(-1), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
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
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v651<<(uint(int32(4))%32)+v644*int32(100))+16)) = v532
	goto L215
L215:
	;
	goto L191
L216:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v40+v668<<(uint(int32(4))%32)+v661*int32(100))+16)) = v532
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
	v732 = int32(1)
	if l4 != 0 {
		v859 = v40
		v865 = v732
		goto L2
	} else {
		goto L225
	}
L222:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	if int32(0) <= v727 {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	F_assign_record_type_typmod(m, v40)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	goto L221
L225:
	;
	v889 = v732
	goto L1
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = l0
	F_errmsg_internal(m, int32(_a_F_internal_get_result_type_2), v29)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L5
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_internal_get_result_type_3), int32(446), int32(_a_F_internal_get_result_type_4))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L5
	} else {
		goto L257
	}
L230:
	;
	if l3 != 0 {
		goto L240
	} else {
		goto L241
	}
L231:
	;
	v769 = F_exprType(m, l1)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L5
	} else {
		goto L238
	}
L232:
	;
	switch v39 - int32(2277) {
	case 0, 6:
		goto L231
	case 1, 2, 3, 4, 5:
		v773 = v39
		goto L230
	default:
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	if base.B2i32(base.Ui32(v39-int32(_a_F_internal_get_result_type_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v39-int32(_a_F_internal_get_result_type_1)) < base.Ui32(int32(2)))|base.B2i32(v39 == int32(3831)) != 0 {
		goto L231
	} else {
		goto L237
	}
L235:
	;
	if base.B2i32(v39 == int32(2776))|base.B2i32(v39 == int32(3500)) != 0 {
		goto L231
	} else {
		goto L236
	}
L236:
	;
	v773 = v39
	goto L230
L237:
	;
	v773 = v39
	goto L230
L238:
	;
	if v769 == int32(0) {
		goto L229
	} else {
		goto L239
	}
L239:
	;
	v773 = v769
	goto L230
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v773
	goto L242
L241:
	;
	goto L242
L242:
	;
	if l4 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L245
L244:
	;
	goto L245
L245:
	;
	v779 = F_get_type_func_class(m, v773, v27+int32(-16))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L5
	} else {
		goto L248
	}
L246:
	;
	v789 = int32(3)
	if l2 == int32(0) {
		v889 = v789
		goto L1
	} else {
		goto L251
	}
L247:
	;
	if l4 == int32(0) {
		v889 = v779
		goto L1
	} else {
		goto L249
	}
L248:
	;
	switch v779 - int32(1) {
	case 0, 1:
		goto L247
	case 2:
		goto L246
	default:
		v889 = v779
		goto L1
	}
L249:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v787 = F_lookup_rowtype_tupdesc_copy(m, v785, int32(-1))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L250
	}
L250:
	;
	v859 = v787
	v865 = v779
	goto L2
L251:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v792 != int32(383) {
		v889 = v789
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v795 = int32(1)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v798 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v799 = v795
	goto L255
L254:
	;
	v799 = int32(3)
	goto L255
L255:
	;
	v800 = int32(0)
	if base.B2i32(l4 == v800)|base.B2i32(v798 == v800) != 0 {
		v889 = v799
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v859 = v798
	v865 = v795
	goto L2
L257:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L258
	}
L258:
	;
	v812 = F_format_type_be(m, v39)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v38 + int32(4)
	F_errmsg(m, int32(_a_F_internal_get_result_type_5), v27+int32(-48))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_internal_get_result_type_3), int32(498), int32(_a_F_internal_get_result_type_4))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	v859 = v854
	v865 = v855
	goto L2
L263:
	;
	m.G0 = v29 - int32(-64)
	return v889
}
func F_internal_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_internal_in_0), int32(373), int32(_a_F_internal_in_1), int32(_a_F_internal_in_2), int32(_a_F_internal_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
